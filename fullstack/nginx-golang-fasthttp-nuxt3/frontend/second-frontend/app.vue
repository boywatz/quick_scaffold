<template>
  <UContainer>
    <UCard class="mt-10">
      <div class="container mx-auto px-4 py-8">
        <h1 class="text-4xl font-bold mb-4">second Frontend</h1>
        <div v-if="health.status === 'OK'" class="text-green-500 mb-4">
          {{ health.message }}
        </div>
        <div v-else-if="health.error" class="text-red-500 mb-4">
          {{ health.error }}
        </div>
        <div v-else class="text-yellow-500 mb-4">
          Checking backend health...
        </div>
      </div>
    </UCard>
  </UContainer>
</template>

<script setup lang="ts">
const health = ref({ status: "", message: "", error: "" });

onMounted(async () => {
  await checkHealth();
});

const checkHealth = async () => {
  const config = useRuntimeConfig();
  const response = await fetch(config.public.API_BASE_URL);
  const data = await response.json();

  health.value = {
    status: data.status,
    message: data.message,
    error: "",
  };

  if (response.status !== 200) {
    health.value = {
      status: "",
      message: "",
      error: "Backend is not available",
    };
  }
};
</script>
