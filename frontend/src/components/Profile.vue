<script setup>
  import { ref, onMounted, computed } from 'vue';

  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const profile = ref(null);
  const error = ref(null);
  const loading = ref(true);

  const nameForm = ref({
    name: '',
  });
  const nameLoading = ref(false);
  const nameSuccess = ref(false);

  const passwordForm = ref({
    currentPassword: '',
    newPassword: '',
    confirmPassword: '',
  });
  const passwordLoading = ref(false);
  const passwordSuccess = ref(false);
  const passwordsDoNotMatch = computed(() => {
    return (
      passwordForm.value.newPassword !== passwordForm.value.confirmPassword &&
      passwordForm.value.confirmPassword !== ''
    );
  });

  onMounted(async () => {
    try {
      const res = await fetch(`${baseURL}/profile`, {
        headers: {
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
      });
      const data = await res.json();
      if (res.ok) {
        profile.value = data;
        nameForm.value.name = data.name || '';
      } else {
        error.value = data.error;
      }
    } catch (err) {
      error.value = err.message;
    } finally {
      loading.value = false;
    }
  });

  const updateName = async () => {
    if (nameLoading.value) return;

    nameLoading.value = true;
    nameSuccess.value = false;
    error.value = null;

    try {
      const res = await fetch(`${baseURL}/profile/${profile.value.id}/name`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
        body: JSON.stringify({
          name: nameForm.value.name,
        }),
      });

      const data = await res.json();

      if (res.ok) {
        profile.value.name = nameForm.value.name;
        nameSuccess.value = true;
        setTimeout(() => {
          nameSuccess.value = false;
        }, 3000);
      } else {
        error.value = data.error;
      }
    } catch (err) {
      error.value = err.message;
    } finally {
      nameLoading.value = false;
    }
  };

  const updatePassword = async () => {
    if (passwordLoading.value) return;

    if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
      return;
    }

    passwordLoading.value = true;
    passwordSuccess.value = false;
    error.value = null;

    try {
      const res = await fetch(`${baseURL}/profile/${profile.value.id}/password`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('token'),
        },
        body: JSON.stringify({
          currentPassword: passwordForm.value.currentPassword,
          newPassword: passwordForm.value.newPassword,
        }),
      });

      const data = await res.json();

      if (res.ok) {
        passwordSuccess.value = true;
        passwordForm.value = {
          currentPassword: '',
          newPassword: '',
          confirmPassword: '',
        };
        setTimeout(() => {
          passwordSuccess.value = false;
        }, 3000);
      } else {
        error.value = data.error;
      }
    } catch (err) {
      error.value = err.message;
    } finally {
      passwordLoading.value = false;
    }
  };
</script>

<template>
  <div class="p-4 max-w-5xl mx-auto">
    <h1 class="text-xl font-bold mb-4 dark:text-white">Profile</h1>
    <div v-if="loading" class="dark:text-white">Loading...</div>
    <div v-else-if="error" class="text-red-500">{{ error }}</div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- Left panel: Profile data -->
      <div class="space-y-6">
        <!-- Base information -->
        <div class="bg-white dark:bg-gray-800 shadow rounded-lg p-4">
          <h2 class="text-lg font-semibold mb-3 dark:text-white">Profile data</h2>
          <p class="mb-2 dark:text-gray-300">
            <strong class="dark:text-white">ID:</strong> {{ profile.id }}
          </p>
          <p class="mb-2 dark:text-gray-300">
            <strong class="dark:text-white">Name:</strong> {{ profile.name }}
          </p>
          <p class="mb-2 dark:text-gray-300">
            <strong class="dark:text-white">Email:</strong> {{ profile.email }}
          </p>
        </div>
      </div>

      <!-- Right content: Forms -->
      <div class="space-y-6">
        <!-- Edit name -->
        <div class="bg-white dark:bg-gray-800 shadow rounded-lg p-4">
          <h2 class="text-lg font-semibold mb-3 dark:text-white">Update name</h2>
          <form @submit.prevent="updateName" class="space-y-3">
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300"
                >Name</label
              >
              <input
                id="name"
                v-model="nameForm.name"
                type="text"
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
                required
              />
            </div>
            <div>
              <button
                type="submit"
                class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800"
                :disabled="nameLoading"
              >
                {{ nameLoading ? 'In progress...' : 'Update name' }}
              </button>
            </div>
            <div v-if="nameSuccess" class="text-green-500 text-sm">Name successfully updated</div>
          </form>
        </div>

        <!-- Change password -->
        <div class="bg-white dark:bg-gray-800 shadow rounded-lg p-4">
          <h2 class="text-lg font-semibold mb-3 dark:text-white">Change password</h2>
          <form @submit.prevent="updatePassword" class="space-y-3">
            <div>
              <label
                for="current-password"
                class="block text-sm font-medium text-gray-700 dark:text-gray-300"
                >Current password</label
              >
              <input
                id="current-password"
                v-model="passwordForm.currentPassword"
                type="password"
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
                required
              />
            </div>
            <div>
              <label
                for="new-password"
                class="block text-sm font-medium text-gray-700 dark:text-gray-300"
                >New password</label
              >
              <input
                id="new-password"
                v-model="passwordForm.newPassword"
                type="password"
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
                required
              />
            </div>
            <div>
              <label
                for="confirm-password"
                class="block text-sm font-medium text-gray-700 dark:text-gray-300"
                >Confirm new password</label
              >
              <input
                id="confirm-password"
                v-model="passwordForm.confirmPassword"
                type="password"
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
                required
              />
            </div>
            <div v-if="passwordsDoNotMatch" class="text-red-500 text-sm">
              Passwords do not match!
            </div>
            <div>
              <button
                type="submit"
                class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800"
                :disabled="passwordLoading"
              >
                {{ passwordLoading ? 'In progress...' : 'Change password' }}
              </button>
            </div>
            <div v-if="passwordSuccess" class="text-green-500 text-sm">
              Password successfully changed
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
