<script setup>
  import { ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const token = ref(localStorage.getItem('token'));
  const isDarkMode = ref(localStorage.getItem('darkMode') === 'true');

  function toggleDarkMode() {
    isDarkMode.value = !isDarkMode.value;
    localStorage.setItem('darkMode', isDarkMode.value.toString());
    updateTheme();
  }

  function updateTheme() {
    if (isDarkMode.value) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }

  function handleAuthenticated() {
    token.value = localStorage.getItem('token');
  }

  function logout() {
    localStorage.removeItem('token');
    token.value = null;
    router.push('/login');
  }

  onMounted(() => {
    updateTheme();

    const currentToken = localStorage.getItem('token');
    if (currentToken) {
      let baseURL = import.meta.env.VITE_API_URL || '/api';
      fetch(`${baseURL}/refresh`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${currentToken}`,
        },
      })
        .then((res) => {
          if (!res.ok) throw new Error('Error refreshing token');
          return res.json();
        })
        .then((data) => {
          localStorage.setItem('token', data.token);
          token.value = data.token;
        })
        .catch(() => {
          logout();
        });
    }
  });
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-50 dark:bg-dark-primary dark:text-gray-100">
    <!-- Navigation bar -->
    <nav class="bg-indigo-600 dark:bg-dark-secondary shadow-lg">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-white text-xl font-bold hover:text-indigo-100"
                >gorque
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <!-- Dark Mode switch -->
            <button
              @click="toggleDarkMode"
              class="mr-4 p-2 rounded-full hover:bg-indigo-700 dark:hover:bg-dark-accent"
            >
              <svg
                v-if="isDarkMode"
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-yellow-300"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
                />
              </svg>
              <svg
                v-else
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
                />
              </svg>
            </button>

            <template v-if="token">
              <div class="hidden md:flex items-center space-x-4">
                <router-link
                  to="/profile"
                  class="text-white hover:text-indigo-100 px-3 py-2 rounded-md text-sm font-medium"
                >
                  <span class="flex items-center">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-5 w-5 mr-1"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                      />
                    </svg>
                    Profile
                  </span>
                </router-link>
                <button
                  @click="logout"
                  class="bg-red-500 hover:bg-red-600 text-white px-3 py-2 rounded-md text-sm font-medium transition duration-150 ease-in-out"
                >
                  <span class="flex items-center">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-5 w-5 mr-1"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                      />
                    </svg>
                    Logout
                  </span>
                </button>
              </div>
            </template>
            <template v-else>
              <div class="hidden md:flex items-center space-x-4">
                <router-link
                  to="/login"
                  class="text-gray-200 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                  >Login
                </router-link>
                <router-link
                  to="/register"
                  class="bg-white text-indigo-600 hover:bg-indigo-50 dark:bg-indigo-700 dark:text-white dark:hover:bg-indigo-800 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
                >
                  Sign up
                </router-link>
              </div>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main content -->
    <router-view
      :token="token"
      @authenticated="handleAuthenticated"
      @logout="logout"
      class="flex-grow"
    />

    <!-- Footer -->
    <footer class="bg-gray-800 text-white dark:bg-gray-900">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex flex-col md:flex-row justify-between items-center">
          <div class="mb-2 md:mb-0">
            <p class="text-sm">
              <span>gorque v0.1.0</span>
              <span class="mx-1">|</span>
              <span>&copy; {{ new Date().getFullYear() }} Ákos Fehér</span>
              <span class="mx-1">|</span>
              <a
                href="https://github.com/aafeher/gorque"
                class="text-indigo-300 hover:text-indigo-100 transition-colors"
                target="_blank"
                rel="noopener noreferrer"
                >GitHub</a
              >
            </p>
          </div>
          <div class="text-sm">
            <span
              >Licensed under
              <a
                href="https://opensource.org/licenses/MIT"
                class="text-indigo-300 hover:text-indigo-100 transition-colors"
                target="_blank"
                rel="noopener noreferrer"
                >MIT License</a
              ></span
            >
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped></style>
