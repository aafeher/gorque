<script setup>
  import { ref, watch, onMounted } from 'vue';
  import { formatDate, formatTime } from '@/utils/time-utils';

  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const props = defineProps({
    selectedDeviceId: {
      type: [String, null],
      required: null,
    },
  });

  const emit = defineEmits(['session-selected']);

  const sessions = ref([]);
  const selectedSession = ref(null);
  const loading = ref(false);
  const error = ref(null);

  const shortenSessionId = (sessionId) => {
    if (!sessionId || sessionId.length <= 12) return sessionId;
    return `${sessionId.substring(0, 6)}...${sessionId.substring(sessionId.length - 6)}`;
  };

  const fetchSessions = async () => {
    if (!props.selectedDeviceId) return;

    loading.value = true;
    error.value = null;

    try {
      const response = await fetch(`${baseURL}/session?device-id=${props.selectedDeviceId}`, {
        headers: {
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`);
      }

      const data = await response.json();
      sessions.value = data.sessions;
    } catch (err) {
      error.value = err.message;
      console.error(err);
    } finally {
      loading.value = false;
    }
  };

  const selectSession = (session) => {
    selectedSession.value = session;
    emit('session-selected', session);
  };

  watch(
    () => props.selectedDeviceId,
    () => {
      selectedSession.value = null;
      fetchSessions();
    }
  );

  onMounted(fetchSessions);
</script>

<template>
  <div class="session-list-container">
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
        <span>Loading routes ...</span>
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

    <div v-else-if="sessions.length === 0" class="empty-list">
      <div
        class="bg-yellow-50 border border-yellow-200 text-yellow-700 dark:bg-yellow-900 dark:border-yellow-700 dark:text-yellow-200 px-4 py-3 rounded relative"
        role="alert"
      >
        <span class="block sm:inline">No available routes for this vehicle.</span>
      </div>
    </div>

    <div v-else>
      <ul class="session-list space-y-2">
        <li
          v-for="session in sessions"
          :key="session.SessionID"
          class="session-item border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-all duration-200 dark:border-gray-700 dark:bg-dark-secondary"
          :class="{
            'selected bg-indigo-50 border-indigo-300 dark:bg-indigo-900 dark:border-indigo-600':
              selectedSession && selectedSession.SessionID === session.SessionID,
          }"
          @click="selectSession(session)"
        >
          <div class="flex items-center p-3">
            <div class="flex-shrink-0 mr-3">
              <div class="bg-indigo-100 p-2 rounded-full dark:bg-indigo-800">&#128338;</div>
            </div>
            <div class="flex-grow">
              <h3 class="text-base font-medium text-gray-800 dark:text-gray-200">
                {{ formatDate(session.StartTime) }} {{ formatTime(session.StartTime) }}-{{
                  formatTime(session.EndTime)
                }}
              </h3>
              <p class="text-xs text-gray-500 dark:text-gray-400" :title="session.SessionID">
                ID: {{ shortenSessionId(session.SessionID) }}
              </p>
              <p v-if="session.Details" class="text-xs text-gray-500 dark:text-gray-400">
                {{ session.Details }}
              </p>
            </div>

            <div
              class="flex-shrink-0 absolute bottom-2 right-2"
              v-if="selectedSession && selectedSession.SessionID === session.SessionID"
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
  .session-list-container {
    width: 100%;
  }

  .session-list {
    list-style: none;
    padding: 0;
  }

  .session-item {
    cursor: pointer;
    position: relative;
  }

  .session-item.selected {
    position: relative;
  }

  .session-item.selected::before {
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
  .error,
  .empty-list {
    padding: 20px;
    text-align: center;
  }
</style>
