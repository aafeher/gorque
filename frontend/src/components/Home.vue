<script setup>
  import { ref, computed } from 'vue';
  import { formatDate, formatTime } from '@/utils/time-utils';
  import DeviceList from './DeviceList.vue';
  import SessionList from './SessionList.vue';
  import DataMap from './DataMap.vue';
  import DataChart from './DataChart.vue';

  const props = defineProps({
    token: {
      type: String,
      required: true,
    },
  });

  const emit = defineEmits(['authenticated', 'logout']);

  const selectedDeviceId = ref(null);
  const selectedSessionId = ref(null);
  const selectedDevice = ref(null);
  const selectedSession = ref(null);

  const activeTab = ref('map');
  const tabs = [
    {
      id: 'vehicle',
      iconClass: 'fas fa-car',
      name: 'Vehicle info',
      activeClass: 'bg-gradient-to-r from-indigo-600 to-blue-500',
    },
    {
      id: 'map',
      iconClass: 'fas fa-route',
      name: 'Route',
      activeClass: 'bg-gradient-to-r from-cyan-500 to-teal-500',
    },
    {
      id: 'chart',
      iconClass: 'fas fa-chart-simple',
      name: 'Charts',
      activeClass: 'bg-gradient-to-r from-green-500 to-emerald-500',
    },
  ];

  const filteredDeviceDetails = computed(() => {
    if (!selectedDevice.value) return {};

    const excludedKeys = [
      'DeviceID',
      'ProfileName',
      'ProfileVehicleType',
      'LicensePlate',
      'EngineType',
      'ManufactureYear',
      'RegisterDate',
      'Status',
    ];
    const details = {};

    Object.keys(selectedDevice.value).forEach((key) => {
      if (!excludedKeys.includes(key)) {
        details[key] = selectedDevice.value[key];
      }
    });

    return details;
  });

  const handleDeviceSelected = (device) => {
    if (selectedDeviceId.value !== device.DeviceID) {
      selectedSessionId.value = null;
    }
    selectedDeviceId.value = device.DeviceID;
    selectedDevice.value = device;
  };

  const handleSessionSelected = (session) => {
    selectedSessionId.value = session.SessionID;
    selectedSession.value = session;
  };

  function handleAuthenticated() {
    emit('authenticated');
  }
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-50 dark:bg-dark-primary">
    <!-- Main content -->
    <main class="flex-grow">
      <div v-if="!token" class="max-w-md mx-auto mt-10 px-4 sm:px-6 lg:px-8">
        <router-view @authenticated="handleAuthenticated" />
      </div>

      <div v-else class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex flex-col lg:flex-row gap-6">
          <!-- Left side panel -->
          <div class="lg:w-1/4 space-y-6">
            <!-- Device list -->
            <div class="bg-white dark:bg-dark-secondary overflow-hidden shadow-sm rounded-lg">
              <div
                class="px-4 py-5 sm:px-6 bg-gradient-to-r from-indigo-600 to-blue-500 dark:from-indigo-800 dark:to-blue-700"
              >
                <h3 class="text-lg font-medium leading-6 text-white">
                  <i class="fas fa-car"></i> Vehicles
                </h3>
              </div>
              <div class="p-4">
                <DeviceList
                  @device-selected="handleDeviceSelected"
                  class="max-h-64 overflow-y-auto"
                />
              </div>
            </div>

            <!-- Session list -->
            <div class="bg-white dark:bg-dark-secondary overflow-hidden shadow-sm rounded-lg">
              <div
                class="px-4 py-5 sm:px-6 bg-gradient-to-r from-blue-500 to-cyan-500 dark:from-blue-700 dark:to-cyan-700"
              >
                <h3 class="text-lg font-medium leading-6 text-white">
                  <i class="fas fa-route"></i> Routes
                </h3>
              </div>
              <div class="p-4">
                <SessionList
                  :selectedDeviceId="selectedDeviceId"
                  @session-selected="handleSessionSelected"
                  class="max-h-64 overflow-y-auto"
                />
              </div>
            </div>
          </div>

          <!-- Right side content -->
          <div class="lg:w-3/4 space-y-6">
            <!-- Tabs -->
            <div class="bg-white dark:bg-dark-secondary overflow-hidden shadow-sm rounded-lg">
              <div class="flex">
                <button
                  v-for="tab in tabs"
                  :key="tab.id"
                  @click="activeTab = tab.id"
                  :class="[
                    'flex-1 py-4 px-4 text-center font-medium transition-all duration-200 text-sm rounded-t-lg',
                    activeTab === tab.id
                      ? 'text-white font-medium ' + tab.activeClass
                      : 'text-gray-600 hover:bg-gray-100 bg-gray-50 dark:text-gray-300 dark:hover:bg-gray-700 dark:bg-gray-800',
                  ]"
                >
                  <i :class="tab.iconClass"></i> {{ tab.name }}
                </button>
              </div>

              <!-- Content container -->
              <div class="p-1">
                <!-- Vehicle info content -->
                <div v-if="activeTab === 'vehicle'" class="p-1">
                  <!-- Device data -->
                  <div
                    v-if="selectedDevice"
                    class="bg-white dark:bg-dark-secondary overflow-hidden shadow-sm rounded-lg"
                  >
                    <div
                      class="px-4 py-5 sm:px-6 bg-gradient-to-r from-indigo-600 to-blue-500 flex justify-between items-center"
                    >
                      <h3 class="text-lg font-medium leading-6 text-white">
                        {{ selectedDevice.ProfileName || 'N/A' }}
                      </h3>
                      <span
                        class="bg-indigo-800 text-white text-xs font-medium px-2.5 py-1 rounded-full"
                      >
                        ID: {{ selectedDevice.DeviceID }}
                      </span>
                    </div>
                    <div class="p-6">
                      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <div
                          v-for="(value, key) in filteredDeviceDetails"
                          :key="key"
                          class="bg-gray-50 dark:bg-dark-primary p-3 rounded-lg"
                        >
                          <p class="text-sm text-gray-500 dark:text-gray-500">{{ key }}</p>
                          <p class="text-md font-medium text-gray-800 dark:text-gray-200">
                            {{ value || 'N/A' }}
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Map content -->
                <div v-if="activeTab === 'map'" class="p-1">
                  <div
                    class="px-4 py-5 sm:px-6 bg-gradient-to-r from-cyan-500 to-teal-500 flex justify-between items-center rounded-t-lg"
                  >
                    <h3 class="text-lg font-medium leading-6 text-white">Map</h3>
                    <span
                      class="bg-cyan-800 text-white text-xs font-medium px-2.5 py-1 rounded-full"
                      v-if="selectedDeviceId && selectedSessionId"
                    >
                      {{ formatDate(selectedSession.StartTime) }}
                      {{ formatTime(selectedSession.StartTime) }}-{{
                        formatTime(selectedSession.EndTime)
                      }}
                    </span>
                    <span
                      class="bg-cyan-800 text-white text-xs font-medium px-2.5 py-1 rounded-full"
                      v-if="selectedDeviceId && selectedSessionId"
                    >
                      ID: {{ selectedSessionId }}
                    </span>
                  </div>
                  <DataMap
                    :selectedDeviceId="selectedDeviceId"
                    :selectedSessionId="selectedSessionId"
                    class="rounded-b-lg overflow-hidden"
                  />
                </div>

                <!-- Charts content -->
                <div v-if="activeTab === 'chart'" class="p-1">
                  <div
                    class="px-4 py-5 sm:px-6 bg-gradient-to-r from-green-500 to-emerald-500 flex justify-between items-center rounded-t-lg"
                  >
                    <h3 class="text-lg font-medium leading-6 text-white">Charts</h3>
                  </div>
                  <div class="p-4 bg-white dark:bg-dark-secondary rounded-b-lg">
                    <DataChart
                      v-if="selectedDeviceId && selectedSessionId"
                      :selectedDeviceId="selectedDeviceId"
                      :selectedSessionId="selectedSessionId"
                    />
                    <p v-else class="text-gray-500 p-4 text-center">
                      Select a vehicle and route to view the chart
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped></style>
