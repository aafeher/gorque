import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

export const useChartStore = defineStore('chart', () => {
  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const chartData = ref({});
  const initialized = ref(false);
  const loading = ref(false);
  const error = ref(null);
  const chartConfigurations = ref([]);
  const currentDeviceId = ref(null);
  const currentSessionId = ref(null);
  const configLoading = ref(false);
  const configError = ref(null);

  function isoToMillis(isoString) {
    return new Date(isoString).getTime();
  }

  function processVariableData(variableKey) {
    if (!chartData.value || Object.keys(chartData.value).length === 0) {
      return [];
    }

    const data = [];

    Object.entries(chartData.value).forEach(([timestamp, item]) => {
      const time = isoToMillis(timestamp);

      if (item[variableKey] !== undefined) {
        data.push({
          x: time,
          y: parseFloat(item[variableKey]),
        });
      }
    });

    return data.sort((a, b) => a.x - b.x);
  }

  const allChartSeries = computed(() => {
    console.log('Calculating allChartSeries with configs:', chartConfigurations.value);

    if (!chartConfigurations.value || chartConfigurations.value.length === 0) {
      console.log('No chart configurations found!');
      return [];
    }

    return chartConfigurations.value.map((chart) => {
      const series = chart.variables.map((variable) => {
        return {
          name: `${variable.name}${variable.unit ? ` (${variable.unit})` : ''}`,
          data: processVariableData(variable.key),
        };
      });

      return {
        id: chart.id,
        title: chart.title,
        type: chart.type,
        yAxisTitle: chart.yAxisTitle,
        series: series,
      };
    });
  });

  const chartSeries = computed(() => {
    if (
      chartConfigurations.value.length === 0 ||
      !chartData.value ||
      Object.keys(chartData.value).length === 0
    ) {
      return getDefaultChartSeries();
    }

    if (allChartSeries.value.length > 0) {
      return allChartSeries.value[0].series;
    }

    return [];
  });

  function getDefaultChartSeries() {
    if (!chartData.value || Object.keys(chartData.value).length === 0) {
      return [];
    }

    const kdData = [];
    const kff1001Data = [];
    const kff1237Data = [];

    Object.entries(chartData.value).forEach(([timestamp, item]) => {
      const time = isoToMillis(timestamp);

      if (item.kd !== undefined) {
        kdData.push({
          x: time,
          y: parseFloat(item.kd),
        });
      }

      if (item.kff1001 !== undefined) {
        kff1001Data.push({
          x: time,
          y: parseFloat(item.kff1001),
        });
      }

      if (item.kff1237 !== undefined) {
        kff1237Data.push({
          x: time,
          y: parseFloat(item.kff1237),
        });
      }
    });

    kdData.sort((a, b) => a.x - b.x);
    kff1001Data.sort((a, b) => a.x - b.x);
    kff1237Data.sort((a, b) => a.x - b.x);

    return [
      {
        name: 'Speed (OBD)',
        data: kdData,
      },
      {
        name: 'Speed (GPS)',
        data: kff1001Data,
      },
      {
        name: 'GPS vs OBD Speed difference',
        data: kff1237Data,
      },
    ];
  }

  function setFallbackConfiguration() {
    chartConfigurations.value = [
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
    console.log('Set fallback configuration:', chartConfigurations.value);
  }

  async function fetchChartConfigurations(deviceId, sessionId) {
    configLoading.value = true;
    configError.value = null;

    try {
      const response = await fetch(`${baseURL}/configuration`, {
        headers: {
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error when fetching chart configurations: ${response.status}`);
      }

      const data = await response.json();
      console.log('Configuration API response:', data);

      if (data && data.configuration && Array.isArray(data.configuration.charts)) {
        chartConfigurations.value = data.configuration.charts;
        console.log('Set chart configurations from API:', chartConfigurations.value);
      } else {
        console.warn('Invalid configuration format, using fallback');
        setFallbackConfiguration();
      }
    } catch (err) {
      configError.value = err.message;
      console.error('Error fetching chart configurations:', err);
      setFallbackConfiguration();
    } finally {
      configLoading.value = false;
    }
  }

  async function fetchChartData(deviceId, sessionId) {
    if (
      initialized.value &&
      currentDeviceId.value === deviceId &&
      currentSessionId.value === sessionId
    ) {
      return;
    }

    loading.value = true;
    error.value = null;

    try {
      await fetchChartConfigurations(deviceId, sessionId);

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
      console.log('Chart data response:', data);

      if (data && data.data) {
        chartData.value = data.data;
        console.log(
          'Chart data set successfully, sample keys:',
          Object.keys(data.data).slice(0, 3)
        );
      } else {
        console.warn('Invalid data format received');
        chartData.value = {};
      }

      currentDeviceId.value = deviceId;
      currentSessionId.value = sessionId;
      initialized.value = true;

      if (!chartConfigurations.value || chartConfigurations.value.length === 0) {
        console.warn('No chart configurations available after data fetch, setting fallback');
        setFallbackConfiguration();
      }
    } catch (err) {
      error.value = err.message;
      console.error('Error fetching chart data:', err);
    } finally {
      loading.value = false;
    }
  }

  setFallbackConfiguration();

  return {
    chartData,
    chartSeries,
    allChartSeries,
    chartConfigurations,
    initialized,
    loading,
    error,
    configLoading,
    configError,
    fetchChartData,
    currentDeviceId,
    currentSessionId,
    setFallbackConfiguration,
  };
});
