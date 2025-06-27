package models

type ProfileDataCode string
type UserDataCode string

//type userDataName string

type UserDataItem struct {
	Obd2      string       `json:"obd2"`
	Http      UserDataCode `json:"http"`
	Unit      string       `json:"unit"`
	ShortName string       `json:"short_name"`
	FullName  string       `json:"full_name"`
}

type UserDataRequest struct {
	Email       string   `form:"eml"`
	V           int      `form:"v"`
	Session     int64    `form:"session"` // Session ID
	ID          string   `form:"id"`      // Device ID
	Time        int64    `form:"time"`    // Timestamp
	Lat         float64  `form:"lat"`
	Lon         float64  `form:"lon"`
	Notice      string   `form:"notice"`
	NoticeClass string   `form:"noticeClass"`
	K03         *float64 `form:"k3" json:"k03,omitempty"`  // Fuel Status
	K04         *float64 `form:"k4" json:"k04,omitempty"`  // Engine Load
	K05         *float64 `form:"k5" json:"k05,omitempty"`  // Engine Coolant Temperature
	K06         *float64 `form:"k6" json:"k06,omitempty"`  // Fuel Trim Bank 1 Short Term
	K07         *float64 `form:"k7" json:"k07,omitempty"`  // Fuel Trim Bank 1 Long Term
	K08         *float64 `form:"k8" json:"k08,omitempty"`  // Fuel Trim Bank 2 Short Term
	K09         *float64 `form:"k9" json:"k09,omitempty"`  // Fuel Trim Bank 2 Long Term
	K0A         *float64 `form:"ka" json:"k0a,omitempty"`  // Fuel pressure
	K0B         *float64 `form:"kb" json:"k0b,omitempty"`  // Intake Manifold Pressure
	K0C         *float64 `form:"kc" json:"k0c,omitempty"`  // Engine RPM
	K0D         *float64 `form:"kd" json:"k0d,omitempty"`  // Speed (OBD)
	K0E         *float64 `form:"ke" json:"k0e,omitempty"`  // Timing Advance
	K0F         *float64 `form:"kf" json:"k0f,omitempty"`  // Intake Air Temperature
	K10         *float64 `form:"k10" json:"k10,omitempty"` // Mass Air Flow Rate
	K11         *float64 `form:"k11" json:"k11,omitempty"` // Throttle Position (Manifold)
	K12         *float64 `form:"k12" json:"k12,omitempty"` // Air Status
	K14         *float64 `form:"k14" json:"k14,omitempty"` // Fuel trim bank 1 sensor 1
	K15         *float64 `form:"k15" json:"k15,omitempty"` // Fuel trim bank 1 sensor 2
	K16         *float64 `form:"k16" json:"k16,omitempty"` // Fuel trim bank 1 sensor 3
	K17         *float64 `form:"k17" json:"k17,omitempty"` // Fuel trim bank 1 sensor 4
	K18         *float64 `form:"k18" json:"k18,omitempty"` // Fuel trim bank 2 sensor 1
	K19         *float64 `form:"k19" json:"k19,omitempty"` // Fuel trim bank 2 sensor 2
	K1A         *float64 `form:"k1a" json:"k1a,omitempty"` // Fuel trim bank 2 sensor 3
	K1B         *float64 `form:"k1b" json:"k1b,omitempty"` // Fuel trim bank 2 sensor 4
	K1F         *float64 `form:"k1f" json:"k1f,omitempty"` // Run time since engine start
	K21         *float64 `form:"k21" json:"k21,omitempty"` // Distance travelled with MIL/CEL lit
	K22         *float64 `form:"k22" json:"k22,omitempty"` // Fuel Rail Pressure (relative to manifold vacuum)
	K23         *float64 `form:"k23" json:"k23,omitempty"` // Fuel Rail Pressure
	K24         *float64 `form:"k24" json:"k24,omitempty"` // O2 Sensor1 Equivalence Ratio
	K25         *float64 `form:"k25" json:"k25,omitempty"` // O2 Sensor2 Equivalence Ratio
	K26         *float64 `form:"k26" json:"k26,omitempty"` // O2 Sensor3 Equivalence Ratio
	K27         *float64 `form:"k27" json:"k27,omitempty"` // O2 Sensor4 Equivalence Ratio
	K28         *float64 `form:"k28" json:"k28,omitempty"` // O2 Sensor5 Equivalence Ratio
	K29         *float64 `form:"k29" json:"k29,omitempty"` // O2 Sensor6 Equivalence Ratio
	K2A         *float64 `form:"k2a" json:"k2a,omitempty"` // O2 Sensor7 Equivalence Ratio
	K2B         *float64 `form:"k2b" json:"k2b,omitempty"` // O2 Sensor8 Equivalence Ratio
	K2C         *float64 `form:"k2c" json:"k2c,omitempty"` // EGR Commanded
	K2D         *float64 `form:"k2d" json:"k2d,omitempty"` // EGR Error
	K2F         *float64 `form:"k2f" json:"k2f,omitempty"` // Fuel Level (From Engine ECU)
	K31         *float64 `form:"k31" json:"k31,omitempty"` // Distance travelled since codes cleared
	K32         *float64 `form:"k32" json:"k32,omitempty"` // Evap System Vapour Pressure
	K33         *float64 `form:"k33" json:"k33,omitempty"` // Barometric pressure (from vehicle)
	K34         *float64 `form:"k34" json:"k34,omitempty"` // O2 Sensor1 Equivalence Ratio (alternate)
	K3C         *float64 `form:"k3c" json:"k3c,omitempty"` // Catalyst Temperature (Bank 1 Sensor 1)
	K3D         *float64 `form:"k3d" json:"k3d,omitempty"` // Catalyst Temperature (Bank 2 Sensor 1)
	K3E         *float64 `form:"k3e" json:"k3e,omitempty"` // Catalyst Temperature (Bank 1 Sensor 2)
	K3F         *float64 `form:"k3f"`                      // Catalyst Temperature (Bank 2 Sensor 2)
	K42         *float64 `form:"k42"`                      // Voltage (Control Module)
	K43         *float64 `form:"k43"`                      // Engine Load (Absolute)
	K44         *float64 `form:"k44"`                      // Commanded Equivalence Ratio (lambda)
	K45         *float64 `form:"k45"`                      // Relative Throttle Position
	K46         *float64 `form:"k46"`                      // Ambient air temp
	K47         *float64 `form:"k47"`                      // Absolute Throttle Position B
	K49         *float64 `form:"k49"`                      // Accelerator Pedal Position D
	K4A         *float64 `form:"k4a"`                      // Accelerator Pedal Position E
	K4B         *float64 `form:"k4b"`                      // Accelerator Pedal Position F
	K52         *float64 `form:"k52"`                      // Ethanol Fuel %
	K5A         *float64 `form:"k5a"`                      // Relative Accelerator Pedal Position
	K5C         *float64 `form:"k5c"`                      // Engine Oil Temperature
	K78         *float64 `form:"k78"`                      // Exhaust Gas Temperature 1
	K79         *float64 `form:"k79"`                      // Exhaust Gas Temperature 2
	KB4         *float64 `form:"kb4"`                      // Transmission Temperature (Method 2)
	KFF1001     *float64 `form:"kff1001"`                  // Speed (GPS)
	KFF1005     *float64 `form:"kff1005"`                  // GPS Longitude
	KFF1006     *float64 `form:"kff1006"`                  // GPS Latitude
	KFF1007     *float64 `form:"kff1007"`
	KFF1010     *float64 `form:"kff1010"` // GPS Altitude
	KFF1201     *float64 `form:"kff1201"` // Miles Per Gallon(Instant)
	KFF1202     *float64 `form:"kff1202"` // Turbo Boost & Vacuum Gauge
	KFF1203     *float64 `form:"kff1203"` // Kilometers Per Litre (Instant)
	KFF1204     *float64 `form:"kff1204"` // Trip Distance
	KFF1205     *float64 `form:"kff1205"` // Trip average MPG
	KFF1206     *float64 `form:"kff1206"` // Trip average KPL
	KFF1207     *float64 `form:"kff1207"` // Litres Per 100 Kilometer (Instant)
	KFF1208     *float64 `form:"kff1208"` // Trip average Litres/100 KM
	KFF120C     *float64 `form:"kff120c"` // Trip distance (stored in vehicle profile)
	KFF1214     *float64 `form:"kff1214"` // O2 Volts Bank 1 sensor 1
	KFF1215     *float64 `form:"kff1215"` // O2 Volts Bank 1 sensor 2
	KFF1216     *float64 `form:"kff1216"` // O2 Volts Bank 1 sensor 3
	KFF1217     *float64 `form:"kff1217"` // O2 Volts Bank 1 sensor 4
	KFF1218     *float64 `form:"kff1218"` // O2 Volts Bank 2 sensor 1
	KFF1219     *float64 `form:"kff1219"` // O2 Volts Bank 2 sensor 2
	KFF121A     *float64 `form:"kff121a"` // O2 Volts Bank 2 sensor 3
	KFF121B     *float64 `form:"kff121b"` // O2 Volts Bank 2 sensor 4
	KFF1220     *float64 `form:"kff1220"` // Acceleration Sensor (X axis)
	KFF1221     *float64 `form:"kff1221"` // Acceleration Sensor (Y axis)
	KFF1222     *float64 `form:"kff1222"` // Acceleration Sensor (Z axis)
	KFF1223     *float64 `form:"kff1223"` // Acceleration Sensor (Total)
	KFF1225     *float64 `form:"kff1225"` // Torque
	KFF1226     *float64 `form:"kff1226"` // Horsepower (At the wheels)
	KFF122D     *float64 `form:"kff122d"` // 0-60mph Time
	KFF122E     *float64 `form:"kff122e"` // 0-100kph Time
	KFF122F     *float64 `form:"kff122f"` // 1/4 mile time
	KFF1230     *float64 `form:"kff1230"` // 1/8 mile time
	KFF1237     *float64 `form:"kff1237"` // GPS vs OBD Speed difference
	KFF1238     *float64 `form:"kff1238"` // Voltage (OBD Adapter)
	KFF1239     *float64 `form:"kff1239"` // GPS Accuracy
	KFF123A     *float64 `form:"kff123a"` // GPS Satellites
	KFF123B     *float64 `form:"kff123b"` // GPS Bearing
	KFF1240     *float64 `form:"kff1240"` // O2 Sensor1 wide-range Voltage
	KFF1241     *float64 `form:"kff1241"` // O2 Sensor2 wide-range Voltage
	KFF1242     *float64 `form:"kff1242"` // O2 Sensor3 wide-range Voltage
	KFF1243     *float64 `form:"kff1243"` // O2 Sensor4 wide-range Voltage
	KFF1244     *float64 `form:"kff1244"` // O2 Sensor5 wide-range Voltage
	KFF1245     *float64 `form:"kff1245"` // O2 Sensor6 wide-range Voltage
	KFF1246     *float64 `form:"kff1246"` // O2 Sensor7 wide-range Voltage
	KFF1247     *float64 `form:"kff1247"` // O2 Sensor8 wide-range Voltage
	KFF1249     *float64 `form:"kff1249"` // Air Fuel Ratio (Measured)
	KFF124A     *float64 `form:"kff124a"` // Tilt (x)
	KFF124B     *float64 `form:"kff124b"` // Tilt (y)
	KFF124C     *float64 `form:"kff124c"` // Tilt (z)
	KFF124D     *float64 `form:"kff124d"` // Air Fuel Ratio (Commanded)
	KFF124F     *float64 `form:"kff124f"` // 0-200kph Time
	KFF1257     *float64 `form:"kff1257"` // CO₂ in g/km (Instantaneous)
	KFF1258     *float64 `form:"kff1258"` // CO₂ in g/km (Average)
	KFF125A     *float64 `form:"kff125a"` // Fuel flow rate/minute
	KFF125C     *float64 `form:"kff125c"` // Fuel cost (trip)
	KFF125D     *float64 `form:"kff125d"` // Fuel flow rate/hour
	KFF125E     *float64 `form:"kff125e"` // 60-120mph Time
	KFF125F     *float64 `form:"kff125f"` // 60-80mph Time
	KFF1260     *float64 `form:"kff1260"` // 40-60mph Time
	KFF1261     *float64 `form:"kff1261"` // 80-100mph Time
	KFF1263     *float64 `form:"kff1263"` // Average trip speed (whilst moving only)
	KFF1264     *float64 `form:"kff1264"` // 100-0kph Time
	KFF1265     *float64 `form:"kff1265"` // 60-0mph Time
	KFF1266     *float64 `form:"kff1266"` // Trip Time (Since journey start)
	KFF1267     *float64 `form:"kff1267"` // Trip time (whilst stationary)
	KFF1268     *float64 `form:"kff1268"` // Trip Time (whilst moving)
	KFF1269     *float64 `form:"kff1269"` // Volumetric Efficiency (Calculated)
	KFF126A     *float64 `form:"kff126a"` // Distance to empty (Estimated)
	KFF126B     *float64 `form:"kff126b"` // Fuel Remaining (Calculated from vehicle profile)
	KFF126D     *float64 `form:"kff126d"` // Cost per mile/km (Instant)
	KFF126E     *float64 `form:"kff126e"` // Cost per mile/km (Trip)
	KFF1270     *float64 `form:"kff1270"` // Barometer (on Android device)
	KFF1271     *float64 `form:"kff1271"` // Fuel used (trip)
	KFF1272     *float64 `form:"kff1272"` // Average trip speed (whilst stopped or moving)
	KFF1273     *float64 `form:"kff1273"` // Engine kW (At the wheels)
	KFF1275     *float64 `form:"kff1275"` // 80-120kph Time
	KFF1276     *float64 `form:"kff1276"` // 60-130mph Time
	KFF1277     *float64 `form:"kff1277"` // 0-30mph Time
	KFF1805     *float64 `form:"kff1805"` // Transmission Temperature (Method 1)
	KFF5201     *float64 `form:"kff5201"` // Miles Per Gallon (Long Term Average)
	KFF5202     *float64 `form:"kff5202"` // Kilometers Per Litre (Long Term Average)
	KFF5203     *float64 `form:"kff5203"` // Litres Per 100 Kilometer (Long Term Average)

	ProfileBoostAdjust  *float64 `form:"profileBoostAdjust"`
	ProfileDisplacement *float64 `form:"profileDisplacement"`
	ProfileDragCoeff    *float64 `form:"profileDragCoeff"`
	ProfileFuelCost     *float64 `form:"profileFuelCost"`
	ProfileFuelType     *int64   `form:"profileFuelType"`
	ProfileMPGAdjust    *float64 `form:"profileMPGAdjust"`
	ProfileName         *string  `form:"profileName"`
	ProfileOBDAdjust    *float64 `form:"profileOBDAdjust"`
	ProfileOdometer     *int64   `form:"profileOdometer"`
	ProfileTankCapacity *float64 `form:"profileTankCapacity"`
	ProfileTankUsed     *float64 `form:"profileTankUsed"`
	ProfileVe           *float64 `form:"profileVe"`
	ProfileVehicleType  *int64   `form:"profileVehicleType"`
	ProfileWeight       *float64 `form:"profileWeight"`

	DefaultUnit04     *string `form:"defaultUnit04"`
	DefaultUnit05     *string `form:"defaultUnit05"`
	DefaultUnit06     *string `form:"defaultUnit06"`
	DefaultUnit07     *string `form:"defaultUnit07"`
	DefaultUnit0b     *string `form:"defaultUnit0b"`
	DefaultUnit0c     *string `form:"defaultUnit0c"`
	DefaultUnit0d     *string `form:"defaultUnit0d"`
	DefaultUnit0e     *string `form:"defaultUnit0e"`
	DefaultUnit0f     *string `form:"defaultUnit0f"`
	DefaultUnit10     *string `form:"defaultUnit10"`
	DefaultUnit11     *string `form:"defaultUnit11"`
	DefaultUnit14     *string `form:"defaultUnit14"`
	DefaultUnit15     *string `form:"defaultUnit15"`
	DefaultUnit1f     *string `form:"defaultUnit1f"`
	DefaultUnit21     *string `form:"defaultUnit21"`
	DefaultUnit31     *string `form:"defaultUnit31"`
	DefaultUnit33     *string `form:"defaultUnit33"`
	DefaultUnit42     *string `form:"defaultUnit42"`
	DefaultUnit43     *string `form:"defaultUnit43"`
	DefaultUnit44     *string `form:"defaultUnit44"`
	DefaultUnit45     *string `form:"defaultUnit45"`
	DefaultUnit46     *string `form:"defaultUnit46"`
	DefaultUnit47     *string `form:"defaultUnit47"`
	DefaultUnit49     *string `form:"defaultUnit49"`
	DefaultUnit4a     *string `form:"defaultUnit4a"`
	DefaultUnitFF1001 *string `form:"defaultUnitff1001"`
	DefaultUnitFF1005 *string `form:"defaultUnitff1005"`
	DefaultUnitFF1006 *string `form:"defaultUnitff1006"`
	DefaultUnitFF1007 *string `form:"defaultUnitff1007"`
	DefaultUnitFF1201 *string `form:"defaultUnitff1201"`
	DefaultUnitFF1202 *string `form:"defaultUnitff1202"`
	DefaultUnitFF1203 *string `form:"defaultUnitff1203"`
	DefaultUnitFF1204 *string `form:"defaultUnitff1204"`
	DefaultUnitFF1206 *string `form:"defaultUnitff1206"`
	DefaultUnitFF1208 *string `form:"defaultUnitff1208"`
	DefaultUnitFF120c *string `form:"defaultUnitff120c"`
	DefaultUnitFF1214 *string `form:"defaultUnitff1214"`
	DefaultUnitFF1215 *string `form:"defaultUnitff1215"`
	DefaultUnitFF1225 *string `form:"defaultUnitff1225"`
	DefaultUnitFF1226 *string `form:"defaultUnitff1226"`
	DefaultUnitFF1237 *string `form:"defaultUnitff1237"`
	DefaultUnitFF1238 *string `form:"defaultUnitff1238"`
	DefaultUnitFF123a *string `form:"defaultUnitff123a"`
	DefaultUnitFF124d *string `form:"defaultUnitff124d"`
	DefaultUnitFF125a *string `form:"defaultUnitff125a"`
	DefaultUnitFF125c *string `form:"defaultUnitff125c"`
	DefaultUnitFF125d *string `form:"defaultUnitff125d"`
	DefaultUnitFF1264 *string `form:"defaultUnitff1264"`
	DefaultUnitFF1265 *string `form:"defaultUnitff1265"`
	DefaultUnitFF1266 *string `form:"defaultUnitff1266"`
	DefaultUnitFF1267 *string `form:"defaultUnitff1267"`
	DefaultUnitFF1268 *string `form:"defaultUnitff1268"`
	DefaultUnitFF126a *string `form:"defaultUnitff126a"`
	DefaultUnitFF126b *string `form:"defaultUnitff126b"`
	DefaultUnitFF1271 *string `form:"defaultUnitff1271"`
	DefaultUnitFF1272 *string `form:"defaultUnitff1272"`
	DefaultUnitFF1273 *string `form:"defaultUnitff1273"`
	DefaultUnitFF5201 *string `form:"defaultUnitff5201"`
	DefaultUnitFF5202 *string `form:"defaultUnitff5202"`
	DefaultUnitFF5203 *string `form:"defaultUnitff5203"`

	UserFullName04      *string `form:"userFullName04"`
	UserFullName05      *string `form:"userFullName05"`
	UserFullName06      *string `form:"userFullName06"`
	UserFullName07      *string `form:"userFullName07"`
	UserFullName0b      *string `form:"userFullName0b"`
	UserFullName0c      *string `form:"userFullName0c"`
	UserFullName0d      *string `form:"userFullName0d"`
	UserFullName0e      *string `form:"userFullName0e"`
	UserFullName0f      *string `form:"userFullName0f"`
	UserFullName10      *string `form:"userFullName10"`
	UserFullName11      *string `form:"userFullName11"`
	UserFullName14      *string `form:"userFullName14"`
	UserFullName15      *string `form:"userFullName15"`
	UserFullName1f      *string `form:"userFullName1f"`
	UserFullName21      *string `form:"userFullName21"`
	UserFullName31      *string `form:"userFullName31"`
	UserFullName33      *string `form:"userFullName33"`
	UserFullName42      *string `form:"userFullName42"`
	UserFullName43      *string `form:"userFullName43"`
	UserFullName44      *string `form:"userFullName44"`
	UserFullName45      *string `form:"userFullName45"`
	UserFullName46      *string `form:"userFullName46"`
	UserFullName47      *string `form:"userFullName47"`
	UserFullName49      *string `form:"userFullName49"`
	UserFullName4a      *string `form:"userFullName4a"`
	UserFullNameFF1201  *string `form:"userFullNameff1201"`
	UserFullNameFF1202  *string `form:"userFullNameff1202"`
	UserFullNameFF1203  *string `form:"userFullNameff1203"`
	UserFullNameFF1204  *string `form:"userFullNameff1204"`
	UserFullNameFF1206  *string `form:"userFullNameff1206"`
	UserFullNameFF1208  *string `form:"userFullNameff1208"`
	UserFullNameFF120c  *string `form:"userFullNameff120c"`
	UserFullNameFF1214  *string `form:"userFullNameff1214"`
	UserFullNameFF1215  *string `form:"userFullNameff1215"`
	UserFullNameFF1225  *string `form:"userFullNameff1225"`
	UserFullNameFF1226  *string `form:"userFullNameff1226"`
	UserFullNameFF1237  *string `form:"userFullNameff1237"`
	UserFullNameFF1238  *string `form:"userFullNameff1238"`
	UserFullNameFF123a  *string `form:"userFullNameff123a"`
	UserFullNameFF124d  *string `form:"userFullNameff124d"`
	UserFullNameFF125a  *string `form:"userFullNameff125a"`
	UserFullNameFF125c  *string `form:"userFullNameff125c"`
	UserFullNameFF125d  *string `form:"userFullNameff125d"`
	UserFullNameFF1264  *string `form:"userFullNameff1264"`
	UserFullNameFF1265  *string `form:"userFullNameff1265"`
	UserFullNameFF1266  *string `form:"userFullNameff1266"`
	UserFullNameFF1267  *string `form:"userFullNameff1267"`
	UserFullNameFF1268  *string `form:"userFullNameff1268"`
	UserFullNameFF126a  *string `form:"userFullNameff126a"`
	UserFullNameFF126b  *string `form:"userFullNameff126b"`
	UserFullNameFF1271  *string `form:"userFullNameff1271"`
	UserFullNameFF1272  *string `form:"userFullNameff1272"`
	UserFullNameFF1273  *string `form:"userFullNameff1273"`
	UserFullNameFF5201  *string `form:"userFullNameff5201"`
	UserFullNameFF5202  *string `form:"userFullNameff5202"`
	UserFullNameFF5203  *string `form:"userFullNameff5203"`
	UserShortName04     *string `form:"userShortName04"`
	UserShortName05     *string `form:"userShortName05"`
	UserShortName06     *string `form:"userShortName06"`
	UserShortName07     *string `form:"userShortName07"`
	UserShortName0b     *string `form:"userShortName0b"`
	UserShortName0c     *string `form:"userShortName0c"`
	UserShortName0d     *string `form:"userShortName0d"`
	UserShortName0e     *string `form:"userShortName0e"`
	UserShortName0f     *string `form:"userShortName0f"`
	UserShortName10     *string `form:"userShortName10"`
	UserShortName11     *string `form:"userShortName11"`
	UserShortName14     *string `form:"userShortName14"`
	UserShortName15     *string `form:"userShortName15"`
	UserShortName1f     *string `form:"userShortName1f"`
	UserShortName21     *string `form:"userShortName21"`
	UserShortName31     *string `form:"userShortName31"`
	UserShortName33     *string `form:"userShortName33"`
	UserShortName42     *string `form:"userShortName42"`
	UserShortName43     *string `form:"userShortName43"`
	UserShortName44     *string `form:"userShortName44"`
	UserShortName45     *string `form:"userShortName45"`
	UserShortName46     *string `form:"userShortName46"`
	UserShortName47     *string `form:"userShortName47"`
	UserShortName49     *string `form:"userShortName49"`
	UserShortName4a     *string `form:"userShortName4a"`
	UserShortNameFF1201 *string `form:"userShortNameff1201"`
	UserShortNameFF1202 *string `form:"userShortNameff1202"`
	UserShortNameFF1203 *string `form:"userShortNameff1203"`
	UserShortNameFF1204 *string `form:"userShortNameff1204"`
	UserShortNameFF1206 *string `form:"userShortNameff1206"`
	UserShortNameFF1208 *string `form:"userShortNameff1208"`
	UserShortNameFF120c *string `form:"userShortNameff120c"`
	UserShortNameFF1214 *string `form:"userShortNameff1214"`
	UserShortNameFF1215 *string `form:"userShortNameff1215"`
	UserShortNameFF1225 *string `form:"userShortNameff1225"`
	UserShortNameFF1226 *string `form:"userShortNameff1226"`
	UserShortNameFF1237 *string `form:"userShortNameff1237"`
	UserShortNameFF1238 *string `form:"userShortNameff1238"`
	UserShortNameFF123a *string `form:"userShortNameff123a"`
	UserShortNameFF124d *string `form:"userShortNameff124d"`
	UserShortNameFF125a *string `form:"userShortNameff125a"`
	UserShortNameFF125c *string `form:"userShortNameff125c"`
	UserShortNameFF125d *string `form:"userShortNameff125d"`
	UserShortNameFF1264 *string `form:"userShortNameff1264"`
	UserShortNameFF1265 *string `form:"userShortNameff1265"`
	UserShortNameFF1266 *string `form:"userShortNameff1266"`
	UserShortNameFF1267 *string `form:"userShortNameff1267"`
	UserShortNameFF1268 *string `form:"userShortNameff1268"`
	UserShortNameFF126a *string `form:"userShortNameff126a"`
	UserShortNameFF126b *string `form:"userShortNameff126b"`
	UserShortNameFF1271 *string `form:"userShortNameff1271"`
	UserShortNameFF1272 *string `form:"userShortNameff1272"`
	UserShortNameFF1273 *string `form:"userShortNameff1273"`
	UserShortNameFF5201 *string `form:"userShortNameff5201"`
	UserShortNameFF5202 *string `form:"userShortNameff5202"`
	UserShortNameFF5203 *string `form:"userShortNameff5203"`
	UserUnit04          *string `form:"userUnit04"`
	UserUnit05          *string `form:"userUnit05"`
	UserUnit06          *string `form:"userUnit06"`
	UserUnit07          *string `form:"userUnit07"`
	UserUnit0b          *string `form:"userUnit0b"`
	UserUnit0c          *string `form:"userUnit0c"`
	UserUnit0d          *string `form:"userUnit0d"`
	UserUnit0e          *string `form:"userUnit0e"`
	UserUnit0f          *string `form:"userUnit0f"`
	UserUnit10          *string `form:"userUnit10"`
	UserUnit11          *string `form:"userUnit11"`
	UserUnit14          *string `form:"userUnit14"`
	UserUnit15          *string `form:"userUnit15"`
	UserUnit1f          *string `form:"userUnit1f"`
	UserUnit21          *string `form:"userUnit21"`
	UserUnit31          *string `form:"userUnit31"`
	UserUnit33          *string `form:"userUnit33"`
	UserUnit42          *string `form:"userUnit42"`
	UserUnit43          *string `form:"userUnit43"`
	UserUnit44          *string `form:"userUnit44"`
	UserUnit45          *string `form:"userUnit45"`
	UserUnit46          *string `form:"userUnit46"`
	UserUnit47          *string `form:"userUnit47"`
	UserUnit49          *string `form:"userUnit49"`
	UserUnit4a          *string `form:"userUnit4a"`
	UserUnitFF1001      *string `form:"userUnitff1001"`
	UserUnitFF1005      *string `form:"userUnitff1005"`
	UserUnitFF1006      *string `form:"userUnitff1006"`
	UserUnitFF1007      *string `form:"userUnitff1007"`
	UserUnitFF1201      *string `form:"userUnitff1201"`
	UserUnitFF1202      *string `form:"userUnitff1202"`
	UserUnitFF1203      *string `form:"userUnitff1203"`
	UserUnitFF1204      *string `form:"userUnitff1204"`
	UserUnitFF1206      *string `form:"userUnitff1206"`
	UserUnitFF1208      *string `form:"userUnitff1208"`
	UserUnitFF120c      *string `form:"userUnitff120c"`
	UserUnitFF1214      *string `form:"userUnitff1214"`
	UserUnitFF1215      *string `form:"userUnitff1215"`
	UserUnitFF1225      *string `form:"userUnitff1225"`
	UserUnitFF1226      *string `form:"userUnitff1226"`
	UserUnitFF1237      *string `form:"userUnitff1237"`
	UserUnitFF1238      *string `form:"userUnitff1238"`
	UserUnitFF123a      *string `form:"userUnitff123a"`
	UserUnitFF124d      *string `form:"userUnitff124d"`
	UserUnitFF125a      *string `form:"userUnitff125a"`
	UserUnitFF125c      *string `form:"userUnitff125c"`
	UserUnitFF125d      *string `form:"userUnitff125d"`
	UserUnitFF1264      *string `form:"userUnitff1264"`
	UserUnitFF1265      *string `form:"userUnitff1265"`
	UserUnitFF1266      *string `form:"userUnitff1266"`
	UserUnitFF1267      *string `form:"userUnitff1267"`
	UserUnitFF1268      *string `form:"userUnitff1268"`
	UserUnitFF126a      *string `form:"userUnitff126a"`
	UserUnitFF126b      *string `form:"userUnitff126b"`
	UserUnitFF1271      *string `form:"userUnitff1271"`
	UserUnitFF1272      *string `form:"userUnitff1272"`
	UserUnitFF1273      *string `form:"userUnitff1273"`
	UserUnitFF5201      *string `form:"userUnitff5201"`
	UserUnitFF5202      *string `form:"userUnitff5202"`
	UserUnitFF5203      *string `form:"userUnitff5203"`
}

var (
	profileDataCodeProfileBoostAdjust  ProfileDataCode = "profileBoostAdjust"
	profileDataCodeProfileDisplacement ProfileDataCode = "profileDisplacement"
	profileDataCodeProfileDragCoeff    ProfileDataCode = "profileDragCoeff"
	profileDataCodeProfileFuelCost     ProfileDataCode = "profileFuelCost"
	profileDataCodeProfileFuelType     ProfileDataCode = "profileFuelType"
	profileDataCodeProfileMPGAdjust    ProfileDataCode = "profileMPGAdjust"
	profileDataCodeProfileName         ProfileDataCode = "profileName"
	profileDataCodeProfileOBDAdjust    ProfileDataCode = "profileOBDAdjust"
	profileDataCodeProfileOdometer     ProfileDataCode = "profileOdometer"
	profileDataCodeProfileTankCapacity ProfileDataCode = "profileTankCapacity"
	profileDataCodeProfileTankUsed     ProfileDataCode = "profileTankUsed"
	profileDataCodeProfileVe           ProfileDataCode = "profileVe"
	profileDataCodeProfileVehicleType  ProfileDataCode = "profileVehicleType"
	profileDataCodeProfileWeight       ProfileDataCode = "profileWeight"

	ProfileDataCodeList = []ProfileDataCode{
		profileDataCodeProfileBoostAdjust,
		profileDataCodeProfileDisplacement,
		profileDataCodeProfileDragCoeff,
		profileDataCodeProfileFuelCost,
		profileDataCodeProfileFuelType,
		profileDataCodeProfileMPGAdjust,
		profileDataCodeProfileName,
		profileDataCodeProfileOBDAdjust,
		profileDataCodeProfileOdometer,
		profileDataCodeProfileTankCapacity,
		profileDataCodeProfileTankUsed,
		profileDataCodeProfileVe,
		profileDataCodeProfileVehicleType,
		profileDataCodeProfileWeight,
	}

	//userDataName04 userDataName = "userFullName04"
	//userDataName05 userDataName = "userFullName05"
	//userDataName06 userDataName = "userFullName06"
	//userDataName07 userDataName = "userFullName07"
	//userDataName0b userDataName = "userFullName0b"
	//userDataName0c userDataName = "userFullName0c"
	//userDataName0d userDataName = "userFullName0d"
	//userDataName0e userDataName = "userFullName0e"
	//userDataName0f userDataName = "userFullName0f"
	//userDataName10 userDataName = "userFullName10"
	//userDataName11 userDataName = "userFullName11"
	//userDataName14 userDataName = "userFullName14"
	//userDataName15 userDataName = "userFullName15"
	//userDataName1f userDataName = "userFullName1f"

	userDataCodeK03     UserDataCode = "k03"
	UserDataCodeK04     UserDataCode = "k04"
	userDataCodeK05     UserDataCode = "k05"
	userDataCodeK06     UserDataCode = "k06"
	userDataCodeK07     UserDataCode = "k07"
	userDataCodeK08     UserDataCode = "k08"
	userDataCodeK09     UserDataCode = "k09"
	userDataCodeK0A     UserDataCode = "k0a"
	userDataCodeK0B     UserDataCode = "k0b"
	userDataCodeK0C     UserDataCode = "k0c"
	UserDataCodeK0D     UserDataCode = "k0d"
	userDataCodeK0E     UserDataCode = "k0e"
	userDataCodeK0F     UserDataCode = "k0f"
	userDataCodeK10     UserDataCode = "k10"
	userDataCodeK11     UserDataCode = "k11"
	userDataCodeK12     UserDataCode = "k12"
	userDataCodeK14     UserDataCode = "k14"
	userDataCodeK15     UserDataCode = "k15"
	userDataCodeK16     UserDataCode = "k16"
	userDataCodeK17     UserDataCode = "k17"
	userDataCodeK18     UserDataCode = "k18"
	userDataCodeK19     UserDataCode = "k19"
	userDataCodeK1A     UserDataCode = "k1a"
	userDataCodeK1B     UserDataCode = "k1b"
	userDataCodeK1F     UserDataCode = "k1f"
	userDataCodeK21     UserDataCode = "k21"
	userDataCodeK22     UserDataCode = "k22"
	userDataCodeK23     UserDataCode = "k23"
	userDataCodeK24     UserDataCode = "k24"
	userDataCodeK25     UserDataCode = "k25"
	userDataCodeK26     UserDataCode = "k26"
	userDataCodeK27     UserDataCode = "k27"
	userDataCodeK28     UserDataCode = "k28"
	userDataCodeK29     UserDataCode = "k29"
	userDataCodeK2A     UserDataCode = "k2a"
	userDataCodeK2B     UserDataCode = "k2b"
	userDataCodeK2C     UserDataCode = "k2c"
	userDataCodeK2D     UserDataCode = "k2d"
	userDataCodeK2F     UserDataCode = "k2f"
	userDataCodeK31     UserDataCode = "k31"
	userDataCodeK32     UserDataCode = "k32"
	userDataCodeK33     UserDataCode = "k33"
	userDataCodeK34     UserDataCode = "k34"
	userDataCodeK3C     UserDataCode = "k3c"
	userDataCodeK3D     UserDataCode = "k3d"
	userDataCodeK3E     UserDataCode = "k3e"
	userDataCodeK3F     UserDataCode = "k3f"
	userDataCodeK42     UserDataCode = "k42"
	UserDataCodeK43     UserDataCode = "k43"
	userDataCodeK44     UserDataCode = "k44"
	userDataCodeK45     UserDataCode = "k45"
	userDataCodeK46     UserDataCode = "k46"
	userDataCodeK47     UserDataCode = "k47"
	userDataCodeK49     UserDataCode = "k49"
	userDataCodeK4A     UserDataCode = "k4a"
	userDataCodeK4B     UserDataCode = "k4b"
	userDataCodeK52     UserDataCode = "k52"
	userDataCodeK5A     UserDataCode = "k5a"
	userDataCodeK5C     UserDataCode = "k5c"
	userDataCodeK78     UserDataCode = "k78"
	userDataCodeK79     UserDataCode = "k79"
	userDataCodeKB4     UserDataCode = "kb4"
	UserDataCodeKFF1001 UserDataCode = "kff1001"
	userDataCodeKFF1005 UserDataCode = "kff1005"
	userDataCodeKFF1006 UserDataCode = "kff1006"
	userDataCodeKFF1007 UserDataCode = "kff1007"
	userDataCodeKFF1010 UserDataCode = "kff1010"
	UserDataCodeKFF1201 UserDataCode = "kff1201"
	userDataCodeKFF1202 UserDataCode = "kff1202"
	UserDataCodeKFF1203 UserDataCode = "kff1203"
	userDataCodeKFF1204 UserDataCode = "kff1204"
	userDataCodeKFF1205 UserDataCode = "kff1205"
	userDataCodeKFF1206 UserDataCode = "kff1206"
	UserDataCodeKFF1207 UserDataCode = "kff1207"
	userDataCodeKFF1208 UserDataCode = "kff1208"
	userDataCodeKFF120C UserDataCode = "kff120c"
	userDataCodeKFF1214 UserDataCode = "kff1214"
	userDataCodeKFF1215 UserDataCode = "kff1215"
	userDataCodeKFF1216 UserDataCode = "kff1216"
	userDataCodeKFF1217 UserDataCode = "kff1217"
	userDataCodeKFF1218 UserDataCode = "kff1218"
	userDataCodeKFF1219 UserDataCode = "kff1219"
	userDataCodeKFF121A UserDataCode = "kff121a"
	userDataCodeKFF121B UserDataCode = "kff121b"
	userDataCodeKFF1220 UserDataCode = "kff1220"
	userDataCodeKFF1221 UserDataCode = "kff1221"
	userDataCodeKFF1222 UserDataCode = "kff1222"
	userDataCodeKFF1223 UserDataCode = "kff1223"
	UserDataCodeKFF1225 UserDataCode = "kff1225"
	UserDataCodeKFF1226 UserDataCode = "kff1226"
	userDataCodeKFF122D UserDataCode = "kff122d"
	userDataCodeKFF122E UserDataCode = "kff122e"
	userDataCodeKFF122F UserDataCode = "kff122f"
	userDataCodeKFF1230 UserDataCode = "kff1230"
	UserDataCodeKFF1237 UserDataCode = "kff1237"
	userDataCodeKFF1238 UserDataCode = "kff1238"
	userDataCodeKFF1239 UserDataCode = "kff1239"
	userDataCodeKFF123A UserDataCode = "kff123a"
	userDataCodeKFF123B UserDataCode = "kff123b"
	userDataCodeKFF1240 UserDataCode = "kff1240"
	userDataCodeKFF1241 UserDataCode = "kff1241"
	userDataCodeKFF1242 UserDataCode = "kff1242"
	userDataCodeKFF1243 UserDataCode = "kff1243"
	userDataCodeKFF1244 UserDataCode = "kff1244"
	userDataCodeKFF1245 UserDataCode = "kff1245"
	userDataCodeKFF1246 UserDataCode = "kff1246"
	userDataCodeKFF1247 UserDataCode = "kff1247"
	userDataCodeKFF1249 UserDataCode = "kff1249"
	userDataCodeKFF124A UserDataCode = "kff124a"
	userDataCodeKFF124B UserDataCode = "kff124b"
	userDataCodeKFF124C UserDataCode = "kff124c"
	userDataCodeKFF124D UserDataCode = "kff124d"
	userDataCodeKFF124F UserDataCode = "kff124f"
	userDataCodeKFF1257 UserDataCode = "kff1257"
	userDataCodeKFF1258 UserDataCode = "kff1258"
	UserDataCodeKFF125A UserDataCode = "kff125a"
	userDataCodeKFF125C UserDataCode = "kff125c"
	UserDataCodeKFF125D UserDataCode = "kff125d"
	userDataCodeKFF125E UserDataCode = "kff125e"
	userDataCodeKFF125F UserDataCode = "kff125f"
	userDataCodeKFF1260 UserDataCode = "kff1260"
	userDataCodeKFF1261 UserDataCode = "kff1261"
	UserDataCodeKFF1263 UserDataCode = "kff1263"
	userDataCodeKFF1264 UserDataCode = "kff1264"
	userDataCodeKFF1265 UserDataCode = "kff1265"
	userDataCodeKFF1266 UserDataCode = "kff1266"
	userDataCodeKFF1267 UserDataCode = "kff1267"
	userDataCodeKFF1268 UserDataCode = "kff1268"
	userDataCodeKFF1269 UserDataCode = "kff1269"
	userDataCodeKFF126A UserDataCode = "kff126a"
	userDataCodeKFF126B UserDataCode = "kff126b"
	userDataCodeKFF126D UserDataCode = "kff126d"
	userDataCodeKFF126E UserDataCode = "kff126e"
	userDataCodeKFF1270 UserDataCode = "kff1270"
	UserDataCodeKFF1271 UserDataCode = "kff1271"
	UserDataCodeKFF1272 UserDataCode = "kff1272"
	UserDataCodeKFF1273 UserDataCode = "kff1273"
	userDataCodeKFF1275 UserDataCode = "kff1275"
	userDataCodeKFF1276 UserDataCode = "kff1276"
	userDataCodeKFF1277 UserDataCode = "kff1277"
	userDataCodeKFF1805 UserDataCode = "kff1805"
	userDataCodeKFF5201 UserDataCode = "kff5201"
	userDataCodeKFF5202 UserDataCode = "kff5202"
	userDataCodeKFF5203 UserDataCode = "kff5203"

	UserDataItems = map[UserDataCode]UserDataItem{
		userDataCodeK03: {
			Obd2:      "03",
			Http:      userDataCodeK03,
			Unit:      "",
			ShortName: "Fuel Status",
			FullName:  "Fuel Status",
		},
		UserDataCodeK04: {
			Obd2:      "04",
			Http:      UserDataCodeK04,
			Unit:      "%",
			ShortName: "Load",
			FullName:  "Engine Load",
		},
		userDataCodeK05: {
			Obd2:      "05",
			Http:      userDataCodeK05,
			Unit:      "°C",
			ShortName: "Coolant",
			FullName:  "Engine Coolant Temperature",
		},
		userDataCodeK06: {
			Obd2:      "06",
			Http:      userDataCodeK06,
			Unit:      "%",
			ShortName: "STFT",
			FullName:  "Fuel Trim Bank 1 Short Term",
		},
		userDataCodeK07: {
			Obd2:      "07",
			Http:      userDataCodeK07,
			Unit:      "%",
			ShortName: "LTFT",
			FullName:  "Fuel Trim Bank 1 Long Term",
		},
		userDataCodeK08: {
			Obd2:      "08",
			Http:      userDataCodeK08,
			Unit:      "%",
			ShortName: "STFT B2",
			FullName:  "Fuel Trim Bank 2 Short Term",
		},
		userDataCodeK09: {
			Obd2:      "09",
			Http:      userDataCodeK09,
			Unit:      "%",
			ShortName: "LTFT B2",
			FullName:  "Fuel Trim Bank 2 Long Term",
		},
		userDataCodeK0A: {
			Obd2:      "0A",
			Http:      userDataCodeK0A,
			Unit:      "kPa",
			ShortName: "Fuel Pressure",
			FullName:  "Fuel pressure",
		},
		userDataCodeK0B: {
			Obd2:      "0B",
			Http:      userDataCodeK0B,
			Unit:      "psi",
			ShortName: "Intake",
			FullName:  "Intake Manifold Pressure",
		},
		userDataCodeK0C: {
			Obd2:      "0C",
			Http:      userDataCodeK0C,
			Unit:      "rpm",
			ShortName: "Revs",
			FullName:  "Engine RPM",
		},
		UserDataCodeK0D: {
			Obd2:      "0D",
			Http:      UserDataCodeK0D,
			Unit:      "km/h",
			ShortName: "Speed",
			FullName:  "Speed (OBD)",
		},
		userDataCodeK0E: {
			Obd2:      "0E",
			Http:      userDataCodeK0E,
			Unit:      "°",
			ShortName: "Timing Adv",
			FullName:  "Timing Advance",
		},
		userDataCodeK0F: {
			Obd2:      "0F",
			Http:      userDataCodeK0F,
			Unit:      "°C",
			ShortName: "Intake",
			FullName:  "Intake Air Temperature",
		},
		userDataCodeK10: {
			Obd2:      "10",
			Http:      userDataCodeK10,
			Unit:      "g/s",
			ShortName: "MAF",
			FullName:  "Mass Air Flow Rate",
		},
		userDataCodeK11: {
			Obd2:      "11",
			Http:      userDataCodeK11,
			Unit:      "%",
			ShortName: "Throttle",
			FullName:  "Throttle Position(Manifold)",
		},
		userDataCodeK12: {
			Obd2:      "12",
			Http:      userDataCodeK12,
			Unit:      "",
			ShortName: "Air Status",
			FullName:  "Air Status",
		},
		userDataCodeK14: {
			Obd2:      "14",
			Http:      userDataCodeK14,
			Unit:      "%",
			ShortName: "F/TB1S1",
			FullName:  "Fuel trim Bank 1 Sensor 1",
		},
		userDataCodeK15: {
			Obd2:      "15",
			Http:      userDataCodeK15,
			Unit:      "%",
			ShortName: "F/TB1S2",
			FullName:  "Fuel trim Bank 1 Sensor 2",
		},
		userDataCodeK16: {
			Obd2:      "16",
			Http:      userDataCodeK16,
			Unit:      "V",
			ShortName: "O2 B1S3",
			FullName:  "Fuel trim bank 1 sensor 3",
		},
		userDataCodeK17: {
			Obd2:      "17",
			Http:      userDataCodeK17,
			Unit:      "V",
			ShortName: "O2 B1S4",
			FullName:  "Fuel trim bank 1 sensor 4",
		},
		userDataCodeK18: {
			Obd2:      "18",
			Http:      userDataCodeK18,
			Unit:      "V",
			ShortName: "O2 B2S1",
			FullName:  "Fuel trim bank 2 sensor 1",
		},
		userDataCodeK19: {
			Obd2:      "19",
			Http:      userDataCodeK19,
			Unit:      "V",
			ShortName: "O2 B2S2",
			FullName:  "Fuel trim bank 2 sensor 2",
		},
		userDataCodeK1A: {
			Obd2:      "1A",
			Http:      userDataCodeK1A,
			Unit:      "V",
			ShortName: "O2 B2S3",
			FullName:  "Fuel trim bank 2 sensor 3",
		},
		userDataCodeK1B: {
			Obd2:      "1B",
			Http:      userDataCodeK1B,
			Unit:      "V",
			ShortName: "O2 B2S4",
			FullName:  "Fuel trim bank 2 sensor 4",
		},
		userDataCodeK1F: {
			Obd2:      "1F",
			Http:      userDataCodeK1F,
			Unit:      "s",
			ShortName: "RunTime",
			FullName:  "Run time since engine start",
		},
		userDataCodeK21: {
			Obd2:      "21",
			Http:      userDataCodeK21,
			Unit:      "km",
			ShortName: "MIL On",
			FullName:  "Distance travelled with MIL/CEL lit",
		},
		userDataCodeK22: {
			Obd2:      "22",
			Http:      userDataCodeK22,
			Unit:      "kPa",
			ShortName: "Fuel Rail Pressure",
			FullName:  "Fuel Rail Pressure (relative to manifold vacuum)",
		},
		userDataCodeK23: {
			Obd2:      "23",
			Http:      userDataCodeK23,
			Unit:      "kPa",
			ShortName: "Fuel Rail Pressure",
			FullName:  "Fuel Rail Pressure",
		},
		userDataCodeK24: {
			Obd2:      "24",
			Http:      userDataCodeK24,
			Unit:      "",
			ShortName: "O2S1 ER",
			FullName:  "O2 Sensor1 Equivalence Ratio",
		},
		userDataCodeK25: {
			Obd2:      "25",
			Http:      userDataCodeK25,
			Unit:      "",
			ShortName: "O2S2 ER",
			FullName:  "O2 Sensor2 Equivalence Ratio",
		},
		userDataCodeK26: {
			Obd2:      "26",
			Http:      userDataCodeK26,
			Unit:      "",
			ShortName: "O2S3 ER",
			FullName:  "O2 Sensor3 Equivalence Ratio",
		},
		userDataCodeK27: {
			Obd2:      "27",
			Http:      userDataCodeK27,
			Unit:      "",
			ShortName: "O2S4 ER",
			FullName:  "O2 Sensor4 Equivalence Ratio",
		},
		userDataCodeK28: {
			Obd2:      "28",
			Http:      userDataCodeK28,
			Unit:      "",
			ShortName: "O2S5 ER",
			FullName:  "O2 Sensor5 Equivalence Ratio",
		},
		userDataCodeK29: {
			Obd2:      "29",
			Http:      userDataCodeK29,
			Unit:      "",
			ShortName: "O2S6 ER",
			FullName:  "O2 Sensor6 Equivalence Ratio",
		},
		userDataCodeK2A: {
			Obd2:      "2A",
			Http:      userDataCodeK2A,
			Unit:      "",
			ShortName: "O2S7 ER",
			FullName:  "O2 Sensor7 Equivalence Ratio",
		},
		userDataCodeK2B: {
			Obd2:      "2B",
			Http:      userDataCodeK2B,
			Unit:      "",
			ShortName: "O2S8 ER",
			FullName:  "O2 Sensor8 Equivalence Ratio",
		},
		userDataCodeK2C: {
			Obd2:      "2C",
			Http:      userDataCodeK2C,
			Unit:      "%",
			ShortName: "EGR Cmd",
			FullName:  "EGR Commanded",
		},
		userDataCodeK2D: {
			Obd2:      "2D",
			Http:      userDataCodeK2D,
			Unit:      "%",
			ShortName: "EGR Error",
			FullName:  "EGR Error",
		},
		userDataCodeK2F: {
			Obd2:      "2F",
			Http:      userDataCodeK2F,
			Unit:      "%",
			ShortName: "Fuel Level",
			FullName:  "Fuel Level (From Engine ECU)",
		},
		userDataCodeK31: {
			Obd2:      "31",
			Http:      userDataCodeK31,
			Unit:      "km",
			ShortName: "MIL Off",
			FullName:  "Distance travelled since codes cleared",
		},
		userDataCodeK32: {
			Obd2:      "32",
			Http:      userDataCodeK32,
			Unit:      "Pa",
			ShortName: "Evap Pressure",
			FullName:  "Evap System Vapour Pressure",
		},
		userDataCodeK33: {
			Obd2:      "33",
			Http:      userDataCodeK33,
			Unit:      "psi",
			ShortName: "Baro",
			FullName:  "Barometric pressure (from vehicle)",
		},
		userDataCodeK34: {
			Obd2:      "34",
			Http:      userDataCodeK34,
			Unit:      "",
			ShortName: "O2S1 ER Alt",
			FullName:  "O2 Sensor1 Equivalence Ratio (alternate)",
		},
		userDataCodeK3C: {
			Obd2:      "3C",
			Http:      userDataCodeK3C,
			Unit:      "°C",
			ShortName: "Cat Temp B1S1",
			FullName:  "Catalyst Temperature (Bank 1 Sensor 1)",
		},
		userDataCodeK3D: {
			Obd2:      "3D",
			Http:      userDataCodeK3D,
			Unit:      "°C",
			ShortName: "Cat Temp B2S1",
			FullName:  "Catalyst Temperature (Bank 2 Sensor 1)",
		},
		userDataCodeK3E: {
			Obd2:      "3E",
			Http:      userDataCodeK3E,
			Unit:      "°C",
			ShortName: "Cat Temp B1S2",
			FullName:  "Catalyst Temperature (Bank 1 Sensor 2)",
		},
		userDataCodeK3F: {
			Obd2:      "3F",
			Http:      userDataCodeK3F,
			Unit:      "°C",
			ShortName: "Cat Temp B2S2",
			FullName:  "Catalyst Temperature (Bank 2 Sensor 2)",
		},
		userDataCodeK42: {
			Obd2:      "42",
			Http:      userDataCodeK42,
			Unit:      "V",
			ShortName: "Volts(CM)",
			FullName:  "Voltage (Control Module)",
		},
		UserDataCodeK43: {
			Obd2:      "43",
			Http:      UserDataCodeK43,
			Unit:      "%",
			ShortName: "Abs Load",
			FullName:  "Engine Load(Absolute)",
		},
		userDataCodeK44: {
			Obd2:      "44",
			Http:      userDataCodeK44,
			Unit:      "",
			ShortName: "COMEQR",
			FullName:  "Commanded Equivalence Ratio(lambda)",
		},
		userDataCodeK45: {
			Obd2:      "45",
			Http:      userDataCodeK45,
			Unit:      "%",
			ShortName: "R THR",
			FullName:  "Relative Throttle Position",
		},
		userDataCodeK46: {
			Obd2:      "46",
			Http:      userDataCodeK46,
			Unit:      "°C",
			ShortName: "Air temp",
			FullName:  "Ambient air temp",
		},
		userDataCodeK47: {
			Obd2:      "47",
			Http:      userDataCodeK47,
			Unit:      "%",
			ShortName: "A THR2",
			FullName:  "Absolute Throttle Position B",
		},
		userDataCodeK49: {
			Obd2:      "49",
			Http:      userDataCodeK49,
			Unit:      "%",
			ShortName: "PedalD",
			FullName:  "Accelerator PedalPosition D",
		},
		userDataCodeK4A: {
			Obd2:      "4A",
			Http:      userDataCodeK4A,
			Unit:      "%",
			ShortName: "PedalE",
			FullName:  "Accelerator PedalPosition E",
		},
		userDataCodeK4B: {
			Obd2:      "4B",
			Http:      userDataCodeK4B,
			Unit:      "%",
			ShortName: "Accel Pedal Pos F",
			FullName:  "Accelerator Pedal Position F",
		},
		userDataCodeK52: {
			Obd2:      "52",
			Http:      userDataCodeK52,
			Unit:      "%",
			ShortName: "Ethanol Fuel",
			FullName:  "Ethanol Fuel %",
		},
		userDataCodeK5A: {
			Obd2:      "5A",
			Http:      userDataCodeK5A,
			Unit:      "%",
			ShortName: "Rel Accel Pedal Pos",
			FullName:  "Relative Accelerator Pedal Position",
		},
		userDataCodeK5C: {
			Obd2:      "5C",
			Http:      userDataCodeK5C,
			Unit:      "°C",
			ShortName: "Oil Temp",
			FullName:  "Engine Oil Temperature",
		},
		userDataCodeK78: {
			Obd2:      "78",
			Http:      userDataCodeK78,
			Unit:      "°C",
			ShortName: "Exhaust Temp 1",
			FullName:  "Exhaust Gas Temperature 1",
		},
		userDataCodeK79: {
			Obd2:      "79",
			Http:      userDataCodeK79,
			Unit:      "°C",
			ShortName: "Exhaust Temp 2",
			FullName:  "Exhaust Gas Temperature 2",
		},
		userDataCodeKB4: {
			Obd2:      "B4",
			Http:      userDataCodeKB4,
			Unit:      "°C",
			ShortName: "Trans Temp 2",
			FullName:  "Transmission Temperature (Method 2)",
		},
		UserDataCodeKFF1001: {
			Obd2:      "FF1001",
			Http:      UserDataCodeKFF1001,
			Unit:      "km/h",
			ShortName: "GPS Speed",
			FullName:  "Speed (GPS)",
		},
		userDataCodeKFF1005: {
			Obd2:      "FF1005",
			Http:      userDataCodeKFF1005,
			Unit:      "°",
			ShortName: "GPS Longitude",
			FullName:  "GPS Longitude",
		},
		userDataCodeKFF1006: {
			Obd2:      "FF1006",
			Http:      userDataCodeKFF1006,
			Unit:      "°",
			ShortName: "GPS Latitude",
			FullName:  "GPS Latitude",
		},
		userDataCodeKFF1007: {
			Obd2:      "FF1007",
			Http:      userDataCodeKFF1007,
			Unit:      "",
			ShortName: "GPS Info",
			FullName:  "GPS Information",
		},
		userDataCodeKFF1010: {
			Obd2:      "FF1010",
			Http:      userDataCodeKFF1010,
			Unit:      "m",
			ShortName: "GPS Altitude",
			FullName:  "GPS Altitude",
		},
		UserDataCodeKFF1201: {
			Obd2:      "FF1201",
			Http:      UserDataCodeKFF1201,
			Unit:      "mpg",
			ShortName: "MPG",
			FullName:  "Miles Per Gallon(Instant)",
		},
		userDataCodeKFF1202: {
			Obd2:      "FF1202",
			Http:      userDataCodeKFF1202,
			Unit:      "psi",
			ShortName: "Boost",
			FullName:  "Turbo Boost & Vacuum Gauge",
		},
		UserDataCodeKFF1203: {
			Obd2:      "FF1203",
			Http:      UserDataCodeKFF1203,
			Unit:      "kpl",
			ShortName: "KPL",
			FullName:  "Kilometers Per Litre (Instant)",
		},
		userDataCodeKFF1204: {
			Obd2:      "FF1204",
			Http:      userDataCodeKFF1204,
			Unit:      "km",
			ShortName: "Trip",
			FullName:  "Trip Distance",
		},
		userDataCodeKFF1205: {
			Obd2:      "FF1205",
			Http:      userDataCodeKFF1205,
			Unit:      "mpg",
			ShortName: "Trip MPG",
			FullName:  "Trip average MPG",
		},
		userDataCodeKFF1206: {
			Obd2:      "FF1206",
			Http:      userDataCodeKFF1206,
			Unit:      "kpl",
			ShortName: "Trip KPL",
			FullName:  "Trip average KPL",
		},
		UserDataCodeKFF1207: {
			Obd2:      "FF1207",
			Http:      UserDataCodeKFF1207,
			Unit:      "l/100km",
			ShortName: "L/100km Instant",
			FullName:  "Litres Per 100 Kilometer (Instant)",
		},
		userDataCodeKFF1208: {
			Obd2:      "FF1208",
			Http:      userDataCodeKFF1208,
			Unit:      "l/100km",
			ShortName: "Trip LPK",
			FullName:  "Trip average Litres/100 KM",
		},
		userDataCodeKFF120C: {
			Obd2:      "FF120C",
			Http:      userDataCodeKFF120C,
			Unit:      "km",
			ShortName: "Odo",
			FullName:  "Trip distance (stored in vehicle profile)",
		},
		userDataCodeKFF1214: {
			Obd2:      "FF1214",
			Http:      userDataCodeKFF1214,
			Unit:      "V",
			ShortName: "O2B1S1V",
			FullName:  "O2 Bank 1 Sensor 1 Voltage",
		},
		userDataCodeKFF1215: {
			Obd2:      "FF1215",
			Http:      userDataCodeKFF1215,
			Unit:      "V",
			ShortName: "O2B1S2V",
			FullName:  "O2 Bank 1 Sensor 2 Voltage",
		},
		userDataCodeKFF1216: {
			Obd2:      "FF1216",
			Http:      userDataCodeKFF1216,
			Unit:      "V",
			ShortName: "O2 Volts B1S3",
			FullName:  "O2 Volts Bank 1 sensor 3",
		},
		userDataCodeKFF1217: {
			Obd2:      "FF1217",
			Http:      userDataCodeKFF1217,
			Unit:      "V",
			ShortName: "O2 Volts B1S4",
			FullName:  "O2 Volts Bank 1 sensor 4",
		},
		userDataCodeKFF1218: {
			Obd2:      "FF1218",
			Http:      userDataCodeKFF1218,
			Unit:      "V",
			ShortName: "O2 Volts B2S1",
			FullName:  "O2 Volts Bank 2 sensor 1",
		},
		userDataCodeKFF1219: {
			Obd2:      "FF1219",
			Http:      userDataCodeKFF1219,
			Unit:      "V",
			ShortName: "O2 Volts B2S2",
			FullName:  "O2 Volts Bank 2 sensor 2",
		},
		userDataCodeKFF121A: {
			Obd2:      "FF121A",
			Http:      userDataCodeKFF121A,
			Unit:      "V",
			ShortName: "O2 Volts B2S3",
			FullName:  "O2 Volts Bank 2 sensor 3",
		},
		userDataCodeKFF121B: {
			Obd2:      "FF121B",
			Http:      userDataCodeKFF121B,
			Unit:      "V",
			ShortName: "O2 Volts B2S4",
			FullName:  "O2 Volts Bank 2 sensor 4",
		},
		userDataCodeKFF1220: {
			Obd2:      "FF1220",
			Http:      userDataCodeKFF1220,
			Unit:      "g",
			ShortName: "Accel X",
			FullName:  "Acceleration Sensor (X axis)",
		},
		userDataCodeKFF1221: {
			Obd2:      "FF1221",
			Http:      userDataCodeKFF1221,
			Unit:      "g",
			ShortName: "Accel Y",
			FullName:  "Acceleration Sensor (Y axis)",
		},
		userDataCodeKFF1222: {
			Obd2:      "FF1222",
			Http:      userDataCodeKFF1222,
			Unit:      "g",
			ShortName: "Accel Z",
			FullName:  "Acceleration Sensor (Z axis)",
		},
		userDataCodeKFF1223: {
			Obd2:      "FF1223",
			Http:      userDataCodeKFF1223,
			Unit:      "g",
			ShortName: "Accel Total",
			FullName:  "Acceleration Sensor (Total)",
		},
		UserDataCodeKFF1225: {
			Obd2:      "FF1225",
			Http:      UserDataCodeKFF1225,
			Unit:      "Nm",
			ShortName: "Torque",
			FullName:  "Torque",
		},
		UserDataCodeKFF1226: {
			Obd2:      "FF1226",
			Http:      UserDataCodeKFF1226,
			Unit:      "hp",
			ShortName: "HP",
			FullName:  "Horsepower (At the wheels)",
		},
		userDataCodeKFF122D: {
			Obd2:      "FF122D",
			Http:      userDataCodeKFF122D,
			Unit:      "s",
			ShortName: "0-60mph Time",
			FullName:  "0-60mph Time",
		},
		userDataCodeKFF122E: {
			Obd2:      "FF122E",
			Http:      userDataCodeKFF122E,
			Unit:      "s",
			ShortName: "0-100kph Time",
			FullName:  "0-100kph Time",
		},
		userDataCodeKFF122F: {
			Obd2:      "FF122F",
			Http:      userDataCodeKFF122F,
			Unit:      "s",
			ShortName: "1/4 Mile Time",
			FullName:  "1/4 mile time",
		},
		userDataCodeKFF1230: {
			Obd2:      "FF1230",
			Http:      userDataCodeKFF1230,
			Unit:      "s",
			ShortName: "1/8 Mile Time",
			FullName:  "1/8 mile time",
		},
		UserDataCodeKFF1237: {
			Obd2:      "FF1237",
			Http:      UserDataCodeKFF1237,
			Unit:      "km/h",
			ShortName: "Spd Diff",
			FullName:  "GPS vs OBD Speed difference",
		},
		userDataCodeKFF1238: {
			Obd2:      "FF1238",
			Http:      userDataCodeKFF1238,
			Unit:      "V",
			ShortName: "Volts(Ad)",
			FullName:  "Voltage (OBD Adapter)",
		},
		userDataCodeKFF1239: {
			Obd2:      "FF1239",
			Http:      userDataCodeKFF1239,
			Unit:      "m",
			ShortName: "GPS Accuracy",
			FullName:  "GPS Accuracy",
		},
		userDataCodeKFF123A: {
			Obd2:      "FF123A",
			Http:      userDataCodeKFF123A,
			Unit:      "",
			ShortName: "GPS Sat",
			FullName:  "GPS Satellites",
		},
		userDataCodeKFF123B: {
			Obd2:      "FF123B",
			Http:      userDataCodeKFF123B,
			Unit:      "°",
			ShortName: "GPS Bearing",
			FullName:  "GPS Bearing",
		},
		userDataCodeKFF1240: {
			Obd2:      "FF1240",
			Http:      userDataCodeKFF1240,
			Unit:      "V",
			ShortName: "O2 WR Volts 1",
			FullName:  "O2 Sensor1 wide-range Voltage",
		},
		userDataCodeKFF1241: {
			Obd2:      "FF1241",
			Http:      userDataCodeKFF1241,
			Unit:      "V",
			ShortName: "O2 WR Volts 2",
			FullName:  "O2 Sensor2 wide-range Voltage",
		},
		userDataCodeKFF1242: {
			Obd2:      "FF1242",
			Http:      userDataCodeKFF1242,
			Unit:      "V",
			ShortName: "O2 WR Volts 3",
			FullName:  "O2 Sensor3 wide-range Voltage",
		},
		userDataCodeKFF1243: {
			Obd2:      "FF1243",
			Http:      userDataCodeKFF1243,
			Unit:      "V",
			ShortName: "O2 WR Volts 4",
			FullName:  "O2 Sensor4 wide-range Voltage",
		},
		userDataCodeKFF1244: {
			Obd2:      "FF1244",
			Http:      userDataCodeKFF1244,
			Unit:      "V",
			ShortName: "O2 WR Volts 5",
			FullName:  "O2 Sensor5 wide-range Voltage",
		},
		userDataCodeKFF1245: {
			Obd2:      "FF1245",
			Http:      userDataCodeKFF1245,
			Unit:      "V",
			ShortName: "O2 WR Volts 6",
			FullName:  "O2 Sensor6 wide-range Voltage",
		},
		userDataCodeKFF1246: {
			Obd2:      "FF1246",
			Http:      userDataCodeKFF1246,
			Unit:      "V",
			ShortName: "O2 WR Volts 7",
			FullName:  "O2 Sensor7 wide-range Voltage",
		},
		userDataCodeKFF1247: {
			Obd2:      "FF1247",
			Http:      userDataCodeKFF1247,
			Unit:      "V",
			ShortName: "O2 WR Volts 8",
			FullName:  "O2 Sensor8 wide-range Voltage",
		},
		userDataCodeKFF1249: {
			Obd2:      "FF1249",
			Http:      userDataCodeKFF1249,
			Unit:      "",
			ShortName: "AFR Measured",
			FullName:  "Air Fuel Ratio (Measured)",
		},
		userDataCodeKFF124A: {
			Obd2:      "FF124A",
			Http:      userDataCodeKFF124A,
			Unit:      "°",
			ShortName: "Tilt X",
			FullName:  "Tilt (x)",
		},
		userDataCodeKFF124B: {
			Obd2:      "FF124B",
			Http:      userDataCodeKFF124B,
			Unit:      "°",
			ShortName: "Tilt Y",
			FullName:  "Tilt (y)",
		},
		userDataCodeKFF124C: {
			Obd2:      "FF124C",
			Http:      userDataCodeKFF124C,
			Unit:      "°",
			ShortName: "Tilt Z",
			FullName:  "Tilt (z)",
		},
		userDataCodeKFF124D: {
			Obd2:      "FF124D",
			Http:      userDataCodeKFF124D,
			Unit:      ":1",
			ShortName: "AFR(c)",
			FullName:  "Air Fuel Ratio(Commanded)",
		},
		userDataCodeKFF124F: {
			Obd2:      "FF124F",
			Http:      userDataCodeKFF124F,
			Unit:      "s",
			ShortName: "0-200kph Time",
			FullName:  "0-200kph Time",
		},
		userDataCodeKFF1257: {
			Obd2:      "FF1257",
			Http:      userDataCodeKFF1257,
			Unit:      "g/km",
			ShortName: "CO₂ Instant",
			FullName:  "CO₂ in g/km (Instantaneous)",
		},
		userDataCodeKFF1258: {
			Obd2:      "FF1258",
			Http:      userDataCodeKFF1258,
			Unit:      "g/km",
			ShortName: "CO₂ Average",
			FullName:  "CO₂ in g/km (Average)",
		},
		UserDataCodeKFF125A: {
			Obd2:      "FF125A",
			Http:      UserDataCodeKFF125A,
			Unit:      "cc/min",
			ShortName: "Fuel Flow",
			FullName:  "Fuel flow rate/minute",
		},
		userDataCodeKFF125C: {
			Obd2:      "FF125C",
			Http:      userDataCodeKFF125C,
			Unit:      "cost",
			ShortName: "Fuel Cost",
			FullName:  "Fuel cost (trip)",
		},
		UserDataCodeKFF125D: {
			Obd2:      "FF125D",
			Http:      UserDataCodeKFF125D,
			Unit:      "l/hr",
			ShortName: "Fuel Flow",
			FullName:  "Fuel flow rate/hour",
		},
		userDataCodeKFF125E: {
			Obd2:      "FF125E",
			Http:      userDataCodeKFF125E,
			Unit:      "s",
			ShortName: "60-120mph Time",
			FullName:  "60-120mph Time",
		},
		userDataCodeKFF125F: {
			Obd2:      "FF125F",
			Http:      userDataCodeKFF125F,
			Unit:      "s",
			ShortName: "60-80mph Time",
			FullName:  "60-80mph Time",
		},
		userDataCodeKFF1260: {
			Obd2:      "FF1260",
			Http:      userDataCodeKFF1260,
			Unit:      "s",
			ShortName: "40-60mph Time",
			FullName:  "40-60mph Time",
		},
		userDataCodeKFF1261: {
			Obd2:      "FF1261",
			Http:      userDataCodeKFF1261,
			Unit:      "s",
			ShortName: "80-100mph Time",
			FullName:  "80-100mph Time",
		},
		UserDataCodeKFF1263: {
			Obd2:      "FF1263",
			Http:      UserDataCodeKFF1263,
			Unit:      "km/h",
			ShortName: "Avg Trip Speed Moving",
			FullName:  "Average trip speed (whilst moving only)",
		},
		userDataCodeKFF1264: {
			Obd2:      "FF1264",
			Http:      userDataCodeKFF1264,
			Unit:      "s",
			ShortName: "100-0mph",
			FullName:  "100-0mph Time",
		},
		userDataCodeKFF1265: {
			Obd2:      "FF1265",
			Http:      userDataCodeKFF1265,
			Unit:      "s",
			ShortName: "60-0mph",
			FullName:  "60-0mph Time",
		},
		userDataCodeKFF1266: {
			Obd2:      "FF1266",
			Http:      userDataCodeKFF1266,
			Unit:      "s",
			ShortName: "Trip Time",
			FullName:  "Trip Time (Since journey start)",
		},
		userDataCodeKFF1267: {
			Obd2:      "FF1267",
			Http:      userDataCodeKFF1267,
			Unit:      "s",
			ShortName: "Stopped",
			FullName:  "Trip time (whilst stationary)",
		},
		userDataCodeKFF1268: {
			Obd2:      "FF1268",
			Http:      userDataCodeKFF1268,
			Unit:      "s",
			ShortName: "Moving",
			FullName:  "Trip Time (whilst moving)",
		},
		userDataCodeKFF1269: {
			Obd2:      "FF1269",
			Http:      userDataCodeKFF1269,
			Unit:      "%",
			ShortName: "VE",
			FullName:  "Volumetric Efficiency (Calculated)",
		},
		userDataCodeKFF126A: {
			Obd2:      "FF126A",
			Http:      userDataCodeKFF126A,
			Unit:      "km",
			ShortName: "Dist Empt.",
			FullName:  "Distance to empty (Estimated)",
		},
		userDataCodeKFF126B: {
			Obd2:      "FF126B",
			Http:      userDataCodeKFF126B,
			Unit:      "%",
			ShortName: "Fuel Rem",
			FullName:  "Fuel Remaining (Calculated from vehicle profile)",
		},
		userDataCodeKFF126D: {
			Obd2:      "FF126D",
			Http:      userDataCodeKFF126D,
			Unit:      "",
			ShortName: "Cost per Mile/km Instant",
			FullName:  "Cost per mile/km (Instant)",
		},
		userDataCodeKFF126E: {
			Obd2:      "FF126E",
			Http:      userDataCodeKFF126E,
			Unit:      "",
			ShortName: "Cost per Mile/km Trip",
			FullName:  "Cost per mile/km (Trip)",
		},
		userDataCodeKFF1270: {
			Obd2:      "FF1270",
			Http:      userDataCodeKFF1270,
			Unit:      "hPa",
			ShortName: "Barometer",
			FullName:  "Barometer (on Android device)",
		},
		UserDataCodeKFF1271: {
			Obd2:      "FF1271",
			Http:      UserDataCodeKFF1271,
			Unit:      "l",
			ShortName: "Fuel Used",
			FullName:  "Fuel used (trip)",
		},
		UserDataCodeKFF1272: {
			Obd2:      "FF1272",
			Http:      UserDataCodeKFF1272,
			Unit:      "km/h",
			ShortName: "Trip Speed",
			FullName:  "Average trip speed(whilst stopped or moving)",
		},
		UserDataCodeKFF1273: {
			Obd2:      "FF1273",
			Http:      UserDataCodeKFF1273,
			Unit:      "kW",
			ShortName: "Power",
			FullName:  "Engine kW (At the wheels)",
		},
		userDataCodeKFF1275: {
			Obd2:      "FF1275",
			Http:      userDataCodeKFF1275,
			Unit:      "s",
			ShortName: "80-120kph Time",
			FullName:  "80-120kph Time",
		},
		userDataCodeKFF1276: {
			Obd2:      "FF1276",
			Http:      userDataCodeKFF1276,
			Unit:      "s",
			ShortName: "60-130mph Time",
			FullName:  "60-130mph Time",
		},
		userDataCodeKFF1277: {
			Obd2:      "FF1277",
			Http:      userDataCodeKFF1277,
			Unit:      "s",
			ShortName: "0-30mph Time",
			FullName:  "0-30mph Time",
		},
		userDataCodeKFF1805: {
			Obd2:      "FF1805",
			Http:      userDataCodeKFF1805,
			Unit:      "°C",
			ShortName: "Trans Temp 1",
			FullName:  "Transmission Temperature (Method 1)",
		},
		userDataCodeKFF5201: {
			Obd2:      "FF5201",
			Http:      userDataCodeKFF5201,
			Unit:      "mpg",
			ShortName: "MPG(avg)",
			FullName:  "Miles Per Gallon(Long Term Average)",
		},
		userDataCodeKFF5202: {
			Obd2:      "FF5202",
			Http:      userDataCodeKFF5202,
			Unit:      "kpl",
			ShortName: "KPL(avg)",
			FullName:  "Kilometers Per Litre (Long Term Average)",
		},
		userDataCodeKFF5203: {
			Obd2:      "FF5203",
			Http:      userDataCodeKFF5203,
			Unit:      "l/100km",
			ShortName: "LPK(avg)",
			FullName:  "Litres Per 100 Kilometer (Long Term Average)",
		},
	}
)
