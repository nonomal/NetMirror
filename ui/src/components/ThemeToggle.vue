<script setup>
defineProps({
  isDark: Boolean
})

defineEmits(['toggle'])
</script>

<template>
  <button
    @click="$emit('toggle')"
    class="relative inline-flex items-center justify-center w-full h-full rounded-xl hover:bg-gray-100 dark:hover:bg-gray-600 transition-all duration-200"
    aria-label="Toggle theme"
  >
    <span class="sr-only">Toggle theme</span>
    <Transition
      enter-active-class="transition-all duration-300 ease-out-back"
      enter-from-class="opacity-0 rotate-90 scale-50"
      enter-to-class="opacity-100 rotate-0 scale-100"
      leave-active-class="transition-all duration-300 ease-in-back"
      leave-from-class="opacity-100 rotate-0 scale-100"
      leave-to-class="opacity-0 -rotate-90 scale-50"
      mode="out-in"
    >
      <!-- Sun Icon -->
      <svg v-if="isDark" class="w-5 h-5 text-amber-500" key="sun" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z" />
      </svg>
      <!-- Moon Icon -->
      <svg v-else class="w-5 h-5 text-gray-700 dark:text-gray-300" key="moon" fill="currentColor" viewBox="0 0 24 24">
        <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd" />
      </svg>
    </Transition>
  </button>
</template>

<style>
.ease-out-back {
  transition-timing-function: cubic-bezier(0.34, 1.56, 0.64, 1);
}
.ease-in-back {
  transition-timing-function: cubic-bezier(0.68, -0.6, 0.32, 1.6);
}
</style>
