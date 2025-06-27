<script setup>
  import { ref, onMounted } from 'vue';

  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const emit = defineEmits(['device-selected']);

  const devices = ref([]);
  const selectedDevice = ref(null);
  const loading = ref(true);
  const error = ref(null);

  const selectDevice = (device) => {
    selectedDevice.value = device;
    emit('device-selected', device);
  };

  const shortenDeviceId = (deviceId) => {
    if (!deviceId || deviceId.length <= 12) return deviceId;
    return `${deviceId.substring(0, 6)}...${deviceId.substring(deviceId.length - 6)}`;
  };

  const fetchDevices = async () => {
    loading.value = true;
    error.value = null;

    try {
      const response = await fetch(`${baseURL}/device`, {
        headers: {
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`);
      }

      const data = await response.json();
      devices.value = data.devices;

      if (devices.value.length === 1) {
        selectDevice(devices.value[0]);
      }
    } catch (err) {
      error.value = err.message || 'Unknown error.';
      console.error(err);
    } finally {
      loading.value = false;
    }
  };

  onMounted(() => {
    fetchDevices();
  });
</script>

<template>
  <div class="device-list-container">
    <div v-if="loading" class="loading">
      <div class="flex justify-center items-center p-4 dark:text-gray-300">
        <svg
          class="animate-spin h-6 w-6 text-indigo-600 dark:text-indigo-400 mr-2"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
        <span>Loading vehicles ...</span>
      </div>
    </div>

    <div v-else-if="error" class="error">
      <div
        class="bg-red-100 border border-red-400 text-red-700 dark:bg-red-900 dark:border-red-700 dark:text-red-200 px-4 py-3 rounded relative"
        role="alert"
      >
        <strong class="font-bold">Error!</strong>
        <span class="block sm:inline"> {{ error }}</span>
      </div>
    </div>

    <div v-else>
      <ul class="device-list space-y-2">
        <li
          v-for="device in devices"
          :key="device.DeviceID"
          class="device-item border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-all duration-200 dark:border-gray-700 dark:bg-dark-secondary"
          :class="{
            'selected bg-indigo-50 border-indigo-300 dark:bg-indigo-900 dark:border-indigo-600':
              selectedDevice && selectedDevice.DeviceID === device.DeviceID,
          }"
          @click="selectDevice(device)"
        >
          <div class="flex items-center p-3">
            <div class="flex-shrink-0 mr-3">
              <div class="bg-indigo-100 p-2 rounded-full dark:bg-indigo-800">&#128663;</div>
            </div>
            <div class="flex-grow">
              <h3 class="text-base font-medium text-gray-800 dark:text-gray-200">
                {{ device.ProfileName || 'N/A' }}
              </h3>
              <p v-if="device.ProfileVehicleType" class="text-sm text-gray-600 dark:text-gray-400">
                {{ device.ProfileVehicleType }}
              </p>
              <p class="text-xs text-gray-500 dark:text-gray-400" :title="device.DeviceID">
                ID: {{ shortenDeviceId(device.DeviceID) }}
              </p>
            </div>
            <div
              class="flex-shrink-0"
              v-if="selectedDevice && selectedDevice.DeviceID === device.DeviceID"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-indigo-600 dark:text-indigo-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
  .device-list-container {
    width: 100%;
  }

  .device-list {
    list-style: none;
    padding: 0;
  }

  .device-item {
    cursor: pointer;
  }

  .device-item.selected {
    position: relative;
  }

  .device-item.selected::before {
    content: '';
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    width: 4px;
    background-color: #4f46e5;
    border-top-left-radius: 0.5rem;
    border-bottom-left-radius: 0.5rem;
  }

  .loading,
  .error {
    padding: 20px;
    text-align: center;
  }
</style>
