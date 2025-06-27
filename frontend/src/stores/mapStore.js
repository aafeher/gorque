import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useMapStore = defineStore('map', () => {
  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const dataMap = ref({ center: [0.0, 0.0], coords: [], data: {} });
  const initialized = ref(false);
  const loading = ref(false);
  const error = ref(null);

  const currentDeviceId = ref(null);
  const currentSessionId = ref(null);

  async function fetchMapData(deviceId, sessionId) {
    if (
      initialized.value &&
      currentDeviceId.value === deviceId &&
      currentSessionId.value === sessionId &&
      Object.keys(dataMap.value.data).length > 0
    ) {
      return;
    }

    loading.value = true;
    error.value = null;

    try {
      const response = await fetch(
        `${baseURL}/data?device-id=${deviceId}&session-id=${sessionId}`,
        {
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token'),
          },
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`);
      }

      const data = await response.json();
      dataMap.value = data;
      currentDeviceId.value = deviceId;
      currentSessionId.value = sessionId;
      initialized.value = true;
    } catch (err) {
      error.value = err.message;
      console.error(err);
    } finally {
      loading.value = false;
    }
  }

  function resetState() {
    dataMap.value = { center: [0.0, 0.0], coords: [], data: {} };
    initialized.value = false;
    error.value = null;
    currentDeviceId.value = null;
    currentSessionId.value = null;
  }

  return {
    dataMap,
    initialized,
    loading,
    error,
    fetchMapData,
    resetState,
    currentDeviceId,
    currentSessionId,
  };
});
