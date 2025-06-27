<script setup>
  import { computed, watch, onMounted, ref } from 'vue';
  import VueApexCharts from 'vue3-apexcharts';
  import { useChartStore } from '@/stores/chartStore';

  const props = defineProps({
    selectedDeviceId: {
      type: String,
      required: false,
      default: '',
    },
    selectedSessionId: {
      type: Number,
      required: false,
      default: null,
    },
    chartId: {
      type: Number,
      required: false,
      default: null,
    },
  });

  const chartStore = useChartStore();
  const isReady = ref(false);

  async function loadData() {
    if (props.selectedDeviceId && props.selectedSessionId) {
      try {
        await chartStore.fetchChartData(props.selectedDeviceId, props.selectedSessionId);
        isReady.value = true;
      } catch (error) {
        console.error('Error loading data:', error);
      }
    }
  }

  // DEBUG: Providing a simple manual configuration example
  function setupDummyCharts() {
    chartStore.chartConfigurations = [
      {
        id: 1,
        title: 'Speeds',
        type: 'line',
        variables: [
          { key: 'kd', name: 'Speed (OBD)', unit: 'km/h' },
          { key: 'kff1001', name: 'Speed (GPS)', unit: 'km/h' },
          { key: 'kff1237', name: 'GPS vs OBD Speed difference', unit: 'km/h' },
        ],
        yAxisTitle: 'Value (km/h)',
      },
    ];
  }

  function getChartOptions(chart) {
    const isDark = document.documentElement.classList.contains('dark');

    return {
      chart: {
        type: chart.type || 'line',
        height: 350,
        toolbar: {
          show: true,
        },
        zoom: {
          enabled: true,
        },
        foreColor: isDark ? '#d1d5db' : '#374151',
      },
      stroke: {
        curve: 'smooth',
        width: 2,
      },
      title: {
        text: chart.title || 'Chart',
        align: 'center',
        style: {
          color: isDark ? '#f3f4f6' : '#111827',
        },
      },
      xaxis: {
        type: 'datetime',
        title: {
          text: 'Time',
          style: {
            color: isDark ? '#e5e7eb' : '#1f2937',
          },
        },
        labels: {
          datetimeUTC: false,
          format: 'yyyy-MM-dd HH:mm:ss',
          style: {
            colors: isDark ? '#d1d5db' : '#4b5563',
          },
        },
      },
      yaxis: {
        title: {
          text: chart.yAxisTitle || 'Value',
          style: {
            color: isDark ? '#e5e7eb' : '#1f2937',
          },
        },
        labels: {
          style: {
            colors: isDark ? '#d1d5db' : '#4b5563',
          },
        },
      },
      tooltip: {
        x: {
          format: 'yyyy-MM-dd HH:mm:ss',
        },
        theme: isDark ? 'dark' : 'light',
      },
      legend: {
        position: 'top',
        labels: {
          colors: isDark ? '#e5e7eb' : '#1f2937',
        },
      },
      dataLabels: {
        enabled: false,
      },
      grid: {
        borderColor: isDark ? '#374151' : '#e5e7eb',
      },
      theme: {
        mode: isDark ? 'dark' : 'light',
      },
    };
  }

  const chartOptions = computed(() => {
    const isDark = document.documentElement.classList.contains('dark');

    return {
      chart: {
        type: 'line',
        height: 350,
        toolbar: {
          show: true,
        },
        zoom: {
          enabled: true,
        },
        foreColor: isDark ? '#d1d5db' : '#374151',
      },
      stroke: {
        curve: 'smooth',
        width: 2,
      },
      title: {
        text: 'Speeds',
        align: 'center',
        style: {
          color: isDark ? '#f3f4f6' : '#111827',
        },
      },
      xaxis: {
        type: 'datetime',
        title: {
          text: 'Time',
          style: {
            color: isDark ? '#e5e7eb' : '#1f2937',
          },
        },
        labels: {
          datetimeUTC: false,
          format: 'yyyy-MM-dd HH:mm:ss',
          style: {
            colors: isDark ? '#d1d5db' : '#4b5563',
          },
        },
      },
      yaxis: {
        title: {
          text: 'Value (km/h)',
          style: {
            color: isDark ? '#e5e7eb' : '#1f2937',
          },
        },
        labels: {
          style: {
            colors: isDark ? '#d1d5db' : '#4b5563',
          },
        },
      },
      tooltip: {
        x: {
          format: 'yyyy-MM-dd HH:mm:ss',
        },
        theme: isDark ? 'dark' : 'light',
      },
      legend: {
        position: 'top',
        labels: {
          colors: isDark ? '#e5e7eb' : '#1f2937',
        },
      },
      dataLabels: {
        enabled: false,
      },
      grid: {
        borderColor: isDark ? '#374151' : '#e5e7eb',
      },
      theme: {
        mode: isDark ? 'dark' : 'light',
      },
    };
  });

  watch(
    () => document.documentElement.classList.contains('dark'),
    () => {
      chartOptions.value = { ...chartOptions.value };
    }
  );

  watch(
    () => [props.selectedDeviceId, props.selectedSessionId],
    async () => {
      isReady.value = false;
      if (props.selectedDeviceId && props.selectedSessionId) {
        await loadData();
      }
    },
    { immediate: true }
  );

  onMounted(async () => {
    if (props.selectedDeviceId && props.selectedSessionId) {
      await loadData();
    }
  });
</script>

<template>
  <div class="relative w-full">
    <div
      v-if="chartStore.loading || chartStore.configLoading"
      class="absolute inset-0 flex items-center justify-center bg-white dark:bg-gray-800 bg-opacity-80 dark:bg-opacity-80 z-10"
    >
      <div class="text-center p-4">
        <div
          class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-solid border-green-500 border-r-transparent align-middle"
        ></div>
        <p class="mt-2 text-gray-700 dark:text-gray-300">Loading data...</p>
      </div>
    </div>

    <div
      v-else-if="chartStore.error || chartStore.configError"
      class="absolute inset-0 flex items-center justify-center bg-white dark:bg-gray-800 z-10"
    >
      <div class="text-center p-4 text-red-500">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-10 w-10 mx-auto text-red-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
        <p class="mt-2">{{ chartStore.error || chartStore.configError }}</p>
      </div>
    </div>

    <div
      v-else-if="!chartStore.initialized || Object.keys(chartStore.chartData).length === 0"
      class="absolute inset-0 flex items-center justify-center bg-white dark:bg-gray-800 z-10"
    >
      <div class="text-center p-4 text-gray-500">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-10 w-10 mx-auto text-gray-400"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"
          />
        </svg>
        <p class="mt-2">No data to display</p>
      </div>
    </div>

    <div
      v-if="
        !chartStore.loading &&
        !chartStore.configLoading &&
        !chartStore.error &&
        !chartStore.configError &&
        chartStore.initialized &&
        Object.keys(chartStore.chartData).length > 0
      "
    >
      <div v-if="chartStore.chartConfigurations.length > 0">
        <div v-for="chart in chartStore.allChartSeries" :key="chart.id" class="mb-8 h-[400px]">
          <VueApexCharts
            width="100%"
            height="100%"
            :type="chart.type"
            :options="getChartOptions(chart)"
            :series="chart.series"
          />
        </div>

        <div v-if="chartStore.allChartSeries.length === 0" class="text-center p-4 text-gray-500">
          <p>
            No charts to display. (configurations available:
            {{ chartStore.chartConfigurations.length }})
          </p>
        </div>
      </div>

      <div v-else class="h-[400px]">
        <VueApexCharts
          width="100%"
          height="100%"
          type="line"
          :options="chartOptions"
          :series="chartStore.chartSeries"
        />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
