<script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';

  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const email = ref('');
  const password = ref('');
  const router = useRouter();
  const emit = defineEmits(['authenticated']);

  async function login() {
    try {
      const res = await fetch(`${baseURL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: email.value, password: password.value }),
      });
      const data = await res.json();
      if (res.ok) {
        localStorage.setItem('token', data.token);
        emit('authenticated');
        await router.push('/');
      } else {
        alert(data.error);
      }
    } catch (err) {
      alert(err.message);
    }
  }
</script>

<template>
  <div class="p-4 max-w-md mx-auto dark:bg-dark-primary">
    <h1 class="text-xl font-bold mb-4 dark:text-white">Login</h1>
    <form @submit.prevent="login">
      <input
        v-model="email"
        type="email"
        placeholder="Email"
        class="input mb-2 w-full dark:bg-dark-secondary dark:border-gray-700 dark:text-white"
      />
      <input
        v-model="password"
        type="password"
        placeholder="Password"
        class="input mb-2 w-full dark:bg-dark-secondary dark:border-gray-700 dark:text-white"
      />
      <button
        type="submit"
        class="btn btn-blue w-full dark:bg-dark-accent dark:hover:bg-indigo-700"
      >
        Login
      </button>
    </form>
    <div class="mt-4 dark:text-gray-300">
      <router-link to="/register" class="dark:text-indigo-300 dark:hover:text-indigo-200"
        >Don't have an account yet? Sign up!
      </router-link>
    </div>
  </div>
</template>

<style scoped>
  .input {
    border: 1px solid #ccc;
    padding: 8px;
    border-radius: 6px;
  }

  .btn {
    padding: 8px 12px;
    border: none;
    border-radius: 6px;
    background-color: #3490dc;
    color: white;
  }
</style>
