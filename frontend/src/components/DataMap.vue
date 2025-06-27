<script setup>
  import { ref, watch, onMounted, computed, onBeforeUnmount, onUnmounted } from 'vue';
  import { useMapStore } from '@/stores/mapStore';
  import maplibregl from 'maplibre-gl';
  import 'maplibre-gl/dist/maplibre-gl.css';

  const props = defineProps({
    selectedDeviceId: {
      type: [String, null],
      default: null,
    },
    selectedSessionId: {
      type: [Number, String],
      required: null,
    },
  });

  const mapStore = useMapStore();

  const mapContainer = ref(null);
  const mapInstance = ref(null);

  watch(
    () => [props.selectedDeviceId, props.selectedSessionId],
    async () => {
      if (props.selectedDeviceId && props.selectedSessionId) {
        await mapStore.fetchMapData(props.selectedDeviceId, props.selectedSessionId);
        if (!mapStore.loading && Object.keys(mapStore.dataMap.data).length > 0) {
          if (mapInstance.value) {
            updateMap();
          } else {
            initMap();
          }
        }
      }
    },
    { immediate: true }
  );

  onMounted(() => {
    if (props.selectedDeviceId && props.selectedSessionId) {
      mapStore.fetchMapData(props.selectedDeviceId, props.selectedSessionId).then(() => {
        if (!mapStore.loading && Object.keys(mapStore.dataMap.data).length > 0) {
          initMap();
        }
      });
    }
  });

  onUnmounted(() => {
    if (mapInstance.value) {
      mapInstance.value.remove();
      mapInstance.value = null;
    }
  });

  function initMap() {
    if (!mapContainer.value) return;

    let initialCenter = [0, 0];
    let initialZoom = 12;

    console.log(mapStore.dataMap.center);

    if (mapStore.dataMap.center && mapStore.dataMap.center.length > 0) {
      initialCenter = mapStore.dataMap.center;
    }

    mapInstance.value = new maplibregl.Map({
      container: mapContainer.value,
      style: {
        version: 8,
        name: 'OSM Base Map',
        sources: {
          'raster-tiles': {
            type: 'raster',
            tiles: ['https://a.tile.openstreetmap.org/{z}/{x}/{y}.png'],
            tileSize: 256,
            attribution:
              '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap contributors</a>',
          },
        },
        layers: [
          {
            id: 'osm-layer',
            type: 'raster',
            source: 'raster-tiles',
            minzoom: 0,
            maxzoom: 19,
          },
        ],
      },
      center: [initialCenter[1], initialCenter[0]],
      zoom: initialZoom,
    });

    mapInstance.value.addControl(new maplibregl.NavigationControl());
    mapInstance.value.addControl(new maplibregl.ScaleControl());

    mapInstance.value.on('load', () => {
      updateMap();
    });
  }

  function updateMap() {
    if (
      !mapInstance.value ||
      !mapStore.dataMap.data ||
      Object.keys(mapStore.dataMap.data).length === 0
    )
      return;

    const geoJsonData = convertDataToGeoJson(mapStore.dataMap.coords);

    if (!geoJsonData.geometry.coordinates.length) {
      console.warn('There are no coordinates to render on the map.');
      return;
    }

    if (mapInstance.value.getSource('route')) {
      mapInstance.value.getSource('route').setData(geoJsonData);
    } else {
      mapInstance.value.addSource('route', {
        type: 'geojson',
        data: geoJsonData,
        lineMetrics: true,
      });

      mapInstance.value.addLayer({
        id: 'route-line',
        type: 'line',
        source: 'route',
        layout: {
          'line-join': 'round',
          'line-cap': 'round',
        },
        paint: {
          'line-color': '#22B3AA',
          'line-width': 5,
          'line-opacity': 0.7,
        },
      });
    }

    addRouteMarkers(geoJsonData.geometry.coordinates);

    fitMapToData(geoJsonData);
  }

  function convertDataToGeoJson(coordinatesLatLng) {
    let coordinatesLngLat = [];

    coordinatesLatLng.forEach((point) => {
      coordinatesLngLat.push([point[1], point[0]]);
    });

    return {
      type: 'Feature',
      properties: {},
      geometry: {
        type: 'LineString',
        coordinates: coordinatesLngLat,
      },
    };
  }

  function addRouteMarkers(coordinates) {
    if (!coordinates || coordinates.length === 0) return;

    removeExistingMarkers();

    new maplibregl.Marker({ color: '#00ff00' })
      .setLngLat(coordinates[0])
      .setPopup(new maplibregl.Popup().setHTML('<h3>Start</h3>'))
      .addTo(mapInstance.value);

    new maplibregl.Marker({ color: '#ff0000' })
      .setLngLat(coordinates[coordinates.length - 1])
      .setPopup(new maplibregl.Popup().setHTML('<h3>End</h3>'))
      .addTo(mapInstance.value);
  }

  function removeExistingMarkers() {
    if (mapContainer.value) {
      const markers = mapContainer.value.querySelectorAll('.maplibregl-marker');
      markers.forEach((marker) => {
        marker.remove();
      });
    }
  }

  function fitMapToData(geoJson) {
    if (
      !mapInstance.value ||
      !geoJson ||
      !geoJson.geometry ||
      !geoJson.geometry.coordinates ||
      geoJson.geometry.coordinates.length === 0
    )
      return;

    const bounds = new maplibregl.LngLatBounds();
    geoJson.geometry.coordinates.forEach((coord) => {
      bounds.extend(coord);
    });

    mapInstance.value.fitBounds(bounds, {
      padding: 50,
      maxZoom: 15,
    });
  }
</script>

<template>
  <div class="relative h-[400px] w-full">
    <div
      v-if="mapStore.loading"
      class="absolute inset-0 flex items-center justify-center bg-white dark:bg-gray-800 bg-opacity-80 dark:bg-opacity-80 z-10"
    >
      <div class="text-center p-4">
        <div
          class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-solid border-blue-500 border-r-transparent align-middle"
        ></div>
        <p class="mt-2 text-gray-700 dark:text-gray-300">Loading map data ...</p>
      </div>
    </div>

    <div
      v-else-if="mapStore.error"
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
        <p class="mt-2">{{ mapStore.error }}</p>
      </div>
    </div>

    <div
      v-else-if="!mapStore.initialized || Object.keys(mapStore.dataMap.data).length === 0"
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
        <p class="mt-2">There are no data to render map.</p>
      </div>
    </div>

    <div
      ref="mapContainer"
      class="absolute inset-0"
      :class="{
        'opacity-0':
          mapStore.loading ||
          mapStore.error ||
          !mapStore.initialized ||
          Object.keys(mapStore.dataMap.data).length === 0,
      }"
    ></div>
  </div>
</template>

<style scoped></style>
