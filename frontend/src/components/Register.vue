<script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';

  const baseURL = import.meta.env.VITE_API_URL || '/api';

  const email = ref('');
  const password = ref('');
  const router = useRouter();

  async function register() {
    try {
      const res = await fetch(`${baseURL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: email.value, password: password.value }),
      });
      const data = await res.json();
      if (res.ok) {
        alert('Registration successful, please log in now.');
        await router.push('/login');
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
    <h1 class="text-xl font-bold mb-4 dark:text-white">Sign up</h1>
    <form @submit.prevent="register">
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
        class="btn btn-green w-full dark:bg-dark-accent dark:hover:bg-indigo-700"
      >
        Sign up
      </button>
    </form>
    <div class="mt-4 dark:text-gray-300">
      <router-link to="/login" class="dark:text-indigo-300 dark:hover:text-indigo-200"
        >Already have an account? Log in!
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
