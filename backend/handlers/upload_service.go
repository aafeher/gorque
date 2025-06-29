package handlers

import (
	"context"
	"fmt"
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type UploadService struct {
	// dependencies here later
}

func NewUploadService() *UploadService {
	return &UploadService{}
}

type UploadRequest struct {
	Data   models.UserDataRequest
	Fields map[string]any
	User   *models.User
}

// ProcessUpload handles the main upload logic
func (s *UploadService) ProcessUpload(c *gin.Context) {
	// Parse and validate request
	request, err := s.parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure session exists
	if err := s.ensureSession(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Determine and handle call type
	callType := s.determineCallType(request.Fields)
	if err := s.handleByType(c, request, callType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "OK!")
}

// parseRequest parses and validates the incoming request
func (s *UploadService) parseRequest(c *gin.Context) (*UploadRequest, error) {
	var data models.UserDataRequest
	if err := c.ShouldBindQuery(&data); err != nil {
		return nil, err
	}

	if data.Email == "" || data.ID == "" || data.Session == 0 {
		return nil, fmt.Errorf("missing required parameters")
	}

	user, err := models.UserGetByEmail(data.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	fields := s.extractFields(data)

	return &UploadRequest{
		Data:   data,
		Fields: fields,
		User:   user,
	}, nil
}

// ensureSession creates or finds the session
func (s *UploadService) ensureSession(request *UploadRequest) error {
	dataTime := time.Unix(request.Data.Time/1000, (request.Data.Time%1000)*int64(time.Millisecond))
	sessionID := strconv.FormatInt(request.Data.Session, 10)

	_, _, err := models.SessionFindOrCreate(
		sessionID,
		request.Data.ID,
		request.User.ID,
		request.Data.V,
		dataTime,
	)

	return err
}

// extractFields extracts all relevant fields from the request data using reflection
func (s *UploadService) extractFields(data models.UserDataRequest) map[string]any {
	fields := map[string]any{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get("form"); tag != "" {
			// Skip core fields that are handled separately
			if s.isCoreField(tag) {
				continue
			}

			fieldValue := v.Field(i)
			if value := s.extractFieldValue(fieldValue); value != nil {
				fields[tag] = value
			}
		}
	}

	return fields
}

// isCoreField checks if a field is a core field that should be skipped
func (s *UploadService) isCoreField(tag string) bool {
	coreFields := []string{"eml", "v", "session", "id", "time", "lat", "lon"}
	for _, coreField := range coreFields {
		if tag == coreField {
			return true
		}
	}
	return false
}

// extractFieldValue extracts the actual value from a reflect.Value
func (s *UploadService) extractFieldValue(fieldValue reflect.Value) any {
	if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
		if fieldValue.Elem().Kind() == reflect.Float64 {
			return fieldValue.Elem().Float()
		} else if fieldValue.Elem().Kind() == reflect.Int64 {
			return fieldValue.Elem().Int()
		} else if fieldValue.Elem().Kind() == reflect.String {
			return fieldValue.Elem().String()
		}
	} else if fieldValue.Kind() == reflect.String && fieldValue.String() != "" {
		return fieldValue.String()
	} else if fieldValue.Kind() == reflect.Int64 && fieldValue.Int() != 0 {
		return fieldValue.Int()
	}
	return nil
}

// determineCallType determines what type of call this is based on the fields
func (s *UploadService) determineCallType(fields map[string]any) string {
	// Check for notice type
	for key := range fields {
		if strings.HasPrefix(key, "notice") || strings.HasPrefix(key, "noticeClass") {
			return "notice"
		}
	}

	// Check for profile type
	for _, field := range models.ProfileDataCodeList {
		if _, exists := fields[string(field)]; exists {
			return "profile"
		}
	}

	// Check for field definitions
	for key := range fields {
		if strings.HasPrefix(key, "userFullName") ||
			strings.HasPrefix(key, "userShortName") ||
			strings.HasPrefix(key, "userUnit") {
			return "fields"
		}
	}

	// Check for default units
	for key := range fields {
		if strings.HasPrefix(key, "defaultUnit") {
			return "defaultunit"
		}
	}

	// Check for actual data
	for key := range fields {
		if strings.HasPrefix(key, "kff") {
			return "data"
		}
	}

	return "unknown"
}

// handleByType routes the request to the appropriate handler based on call type
func (s *UploadService) handleByType(c *gin.Context, request *UploadRequest, callType string) error {
	switch callType {
	case "notice":
		return s.handleNoticeData(request)
	case "profile":
		return s.handleProfileData(request)
	case "defaultunit":
		return s.handleDefaultUnits(request)
	case "fields":
		return s.handleFieldDefinitions(request)
	case "data":
		return s.handleActualData(c, request)
	default:
		return fmt.Errorf("unknown call type: %s", callType)
	}
}

// handleNoticeData processes notice-related data
func (s *UploadService) handleNoticeData(request *UploadRequest) error {
	_, _, err := models.DeviceFindOrCreate(request.Data.ID, request.User.ID)
	if err != nil {
		log.Printf("Device creation error: %v", err)
		return err
	}
	return nil
}

// handleProfileData processes profile-related data and updates device information
func (s *UploadService) handleProfileData(request *UploadRequest) error {
	profile := models.Device{
		DeviceID: request.Data.ID,
		UserID:   request.User.ID,
	}

	// Map fields to device profile using reflection
	if err := s.mapFieldsToProfile(&profile, request.Fields); err != nil {
		return err
	}

	if err := models.DeviceCreateOrUpdate(&profile); err != nil {
		log.Printf("Profile creation/update error: %v", err)
		return err
	}

	// Update session with vehicle profile
	//sessionID := strconv.FormatInt(request.Data.Session, 10)
	//if err := models.SessionUpdateVehicleProfile(sessionID, profile.ID); err != nil {
	//	log.Printf("Session update error: %v", err)
	//	return err
	//}

	return nil
}

// mapFieldsToProfile maps request fields to device profile using reflection
func (s *UploadService) mapFieldsToProfile(profile *models.Device, fields map[string]any) error {
	t := reflect.TypeOf(*profile)
	v := reflect.ValueOf(profile).Elem()

	for key, value := range fields {
		if err := s.setFieldValue(t, v, key, value); err != nil {
			log.Printf("Error setting field %s: %v", key, err)
			// Continue with other fields even if one fails
		}
	}
	return nil
}

// setFieldValue sets a specific field value using reflection
func (s *UploadService) setFieldValue(t reflect.Type, v reflect.Value, key string, value any) error {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get("gorm"); tag != "" {
			if strings.Contains(tag, "column:"+key) {
				fieldValue := v.Field(i)
				if fieldValue.CanSet() {
					return s.assignValue(fieldValue, value)
				}
			}
		}
	}
	return nil
}

// assignValue assigns a value to a reflect.Value based on its type
func (s *UploadService) assignValue(fieldValue reflect.Value, value any) error {
	switch fieldValue.Kind() {
	case reflect.Float64:
		if floatVal, ok := value.(float64); ok {
			fieldValue.SetFloat(floatVal)
		}
	case reflect.Int64:
		if intVal, ok := value.(int64); ok {
			fieldValue.SetInt(intVal)
		}
	case reflect.String:
		if stringVal, ok := value.(string); ok {
			fieldValue.SetString(stringVal)
		}
	default:
		panic("unhandled default case")
	}
	return nil
}

// handleDefaultUnits processes default unit definitions
func (s *UploadService) handleDefaultUnits(request *UploadRequest) error {
	sessionID := strconv.FormatInt(request.Data.Session, 10)

	for key, value := range request.Fields {
		if !strings.HasPrefix(key, "defaultUnit") {
			continue
		}

		fieldKey := strings.TrimPrefix(key, "defaultUnit")
		if err := s.processDefaultUnit(sessionID, request.User.ID, fieldKey, value); err != nil {
			log.Printf("Error processing default unit for %s: %v", fieldKey, err)
			// Continue with other fields
		}
	}
	return nil
}

// processDefaultUnit processes a single default unit field
func (s *UploadService) processDefaultUnit(sessionID string, userID uint, fieldKey string, value any) error {
	valueStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type for default unit")
	}

	_, created, err := models.SessionFieldFindOrCreate(sessionID, userID, fieldKey, valueStr)
	if err != nil {
		return err
	}

	if !created {
		err = models.SessionFieldUpdateDefaultUnit(sessionID, fieldKey, valueStr)
		if err != nil {
			return err
		}
	}

	return nil
}

// handleFieldDefinitions processes field definition data
func (s *UploadService) handleFieldDefinitions(request *UploadRequest) error {
	sessionID := strconv.FormatInt(request.Data.Session, 10)
	fieldKeysMap := make(map[string]models.SessionField)

	// Group fields by base key
	for key, value := range request.Fields {
		baseKey, fieldType := s.parseFieldDefinitionKey(key)
		if baseKey == "" {
			continue
		}

		if _, exists := fieldKeysMap[baseKey]; !exists {
			fieldKeysMap[baseKey] = models.SessionField{
				SessionID: sessionID,
				UserID:    request.User.ID,
				FieldKey:  baseKey,
			}
		}

		field := fieldKeysMap[baseKey]
		if err := s.setSessionFieldValue(&field, fieldType, value); err != nil {
			log.Printf("Error setting field %s: %v", key, err)
			continue
		}
		fieldKeysMap[baseKey] = field
	}

	// Save all fields
	for _, field := range fieldKeysMap {
		if err := models.SessionFieldSaveOrUpdate(field); err != nil {
			log.Printf("Error saving session field: %v", err)
		}
	}

	return nil
}

// parseFieldDefinitionKey parses a field definition key and returns base key and field type
func (s *UploadService) parseFieldDefinitionKey(key string) (string, string) {
	if strings.HasPrefix(key, "userShortName") {
		return strings.TrimPrefix(key, "userShortName"), "short_name"
	} else if strings.HasPrefix(key, "userFullName") {
		return strings.TrimPrefix(key, "userFullName"), "full_name"
	} else if strings.HasPrefix(key, "userUnit") {
		return strings.TrimPrefix(key, "userUnit"), "unit"
	}
	return "", ""
}

// setSessionFieldValue sets a value on a session field based on field type
func (s *UploadService) setSessionFieldValue(field *models.SessionField, fieldType string, value any) error {
	valueStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type for session field")
	}

	switch fieldType {
	case "short_name":
		field.ShortName = valueStr
	case "full_name":
		field.FullName = valueStr
	case "unit":
		field.Unit = valueStr
	}

	return nil
}

// handleActualData processes actual sensor data and writes to InfluxDB
func (s *UploadService) handleActualData(c *gin.Context, request *UploadRequest) error {
	dataFields := s.extractDataFields(request.Fields)
	if len(dataFields) == 0 {
		return nil
	}

	dataTime := time.Unix(request.Data.Time/1000, (request.Data.Time%1000)*int64(time.Millisecond))
	sessionID := strconv.FormatInt(request.Data.Session, 10)

	// Create InfluxDB point
	point := influxdb2.NewPoint(
		"gorque_data",
		map[string]string{
			"id":      request.Data.ID,
			"session": sessionID,
			"v":       strconv.FormatInt(int64(request.Data.V), 10),
			"eml":     request.Data.Email,
		},
		dataFields,
		dataTime,
	)

	// Write to InfluxDB
	if err := models.DBInfluxWriteAPI.WritePoint(context.Background(), point); err != nil {
		log.Printf("Influx write error: %v", err)
		return err
	}

	if err := models.DBInfluxWriteAPI.Flush(context.Background()); err != nil {
		log.Printf("Influx flush error: %v", err)
		return err
	}

	// Update session activity
	if err := models.SessionUpdateActivityAndRecords(sessionID, dataTime); err != nil {
		log.Printf("Session update error: %v", err)
		return err
	}

	return nil
}

// extractDataFields extracts only the data fields (those starting with 'k')
func (s *UploadService) extractDataFields(fields map[string]any) map[string]any {
	dataFields := map[string]any{}
	for key, value := range fields {
		if strings.HasPrefix(key, "k") {
			dataFields[key] = value
		}
	}
	return dataFields
}
