<template>
  <!-- Edit Node Modal (Agent mode - edit only) -->
  <transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="opacity-0 scale-95"
    enter-to-class="opacity-100 scale-100"
    leave-active-class="transition-all duration-200 ease-in"
    leave-from-class="opacity-100 scale-100"
    leave-to-class="opacity-0 scale-95"
  >
    <div class="fixed inset-0 bg-black/50 backdrop-blur-sm overflow-y-auto h-full w-full z-50 flex items-center justify-center p-4">
      <div class="relative bg-white/95 dark:bg-gray-800/95 backdrop-blur-lg border border-primary-200/30 dark:border-primary-700/30 shadow-2xl rounded-2xl max-w-lg w-full animate-scale-in">
        <!-- Modal Header -->
        <div class="px-8 py-6 border-b border-gray-200/50 dark:border-gray-600/50">
          <div class="flex items-center space-x-4">
            <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-primary-600 rounded-xl shadow-lg shadow-primary-500/25 flex items-center justify-center animate-scale-in">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold bg-gradient-to-r from-gray-900 to-gray-600 dark:from-gray-100 dark:to-gray-300 bg-clip-text text-transparent">
                Edit Node
              </h3>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">
                Update node name and location
              </p>
            </div>
          </div>
        </div>

        <!-- Modal Body -->
        <div class="px-8 py-6">
          <form @submit.prevent="handleSubmit" class="space-y-6">
            <!-- Node Name Field -->
            <div class="animate-slide-up" style="animation-delay: 0.1s;">
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3">
                Node Name *
              </label>
              <div class="relative">
                <input
                  v-model="formData.name"
                  type="text"
                  required
                  placeholder="e.g., London Node"
                  class="w-full px-4 py-3 pl-12 border border-gray-300/50 dark:border-gray-600/50 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500 bg-white/80 dark:bg-gray-700/80 backdrop-blur-sm text-gray-900 dark:text-gray-100 transition-all duration-200 placeholder-gray-500 dark:placeholder-gray-400"
                >
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"></path>
                  </svg>
                </div>
              </div>
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-2 ml-1">
                Display name for this node
              </p>
            </div>

            <!-- Location Field -->
            <div class="animate-slide-up" style="animation-delay: 0.2s;">
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3">
                Location *
              </label>
              <div class="relative">
                <input
                  v-model="formData.location"
                  type="text"
                  required
                  placeholder="e.g., London, UK"
                  class="w-full px-4 py-3 pl-12 border border-gray-300/50 dark:border-gray-600/50 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500 bg-white/80 dark:bg-gray-700/80 backdrop-blur-sm text-gray-900 dark:text-gray-100 transition-all duration-200 placeholder-gray-500 dark:placeholder-gray-400"
                >
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  </svg>
                </div>
              </div>
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-2 ml-1">
                Geographic location of the node
              </p>
            </div>

            <!-- URL Field (Read-only) -->
            <div class="animate-slide-up" style="animation-delay: 0.3s;">
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3">
                Endpoint URL
                <span class="text-xs font-normal text-gray-500 dark:text-gray-400 ml-2">(Read-only)</span>
              </label>
              <div class="relative">
                <div class="w-full px-4 py-3 pl-12 border border-gray-200/50 dark:border-gray-600/50 rounded-xl bg-gray-100/80 dark:bg-gray-700/50 backdrop-blur-sm text-gray-600 dark:text-gray-400 font-mono text-sm break-all">
                  {{ formData.url }}
                </div>
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                  </svg>
                </div>
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                  </svg>
                </div>
              </div>
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-2 ml-1">
                URL is set during agent registration and cannot be changed
              </p>
            </div>

            <!-- Form Actions -->
            <div class="flex justify-end space-x-3 pt-6 animate-slide-up" style="animation-delay: 0.4s;">
              <button
                type="button"
                @click="$emit('close')"
                class="px-6 py-3 bg-white/80 dark:bg-gray-700/80 backdrop-blur-sm text-gray-800 dark:text-gray-200 border border-gray-300/50 dark:border-gray-600/50 rounded-xl hover:bg-white dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500 transition-all duration-200 hover:scale-105 shadow-lg"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="!isFormValid || saving"
                class="px-6 py-3 bg-gradient-to-r from-primary-600 to-primary-700 text-white rounded-xl hover:from-primary-700 hover:to-primary-800 focus:outline-none focus:ring-2 focus:ring-primary-500 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 transform hover:scale-105 shadow-lg shadow-primary-500/25"
              >
                <span v-if="saving" class="flex items-center">
                  <svg class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Saving...
                </span>
                <span v-else class="flex items-center">
                  <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                  Save Changes
                </span>
              </button>
            </div>
          </form>
        </div>

        <!-- Info Box -->
        <div class="px-8 pb-6 animate-slide-up" style="animation-delay: 0.5s;">
          <div class="p-4 bg-gradient-to-r from-blue-50 to-primary-50 dark:from-blue-900/20 dark:to-primary-900/20 border border-blue-200/50 dark:border-blue-700/50 rounded-xl backdrop-blur-sm">
            <div class="flex items-start space-x-3">
              <div class="flex-shrink-0">
                <svg class="w-5 h-5 text-blue-600 dark:text-blue-400 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div>
                <h4 class="text-sm font-semibold text-blue-800 dark:text-blue-300 mb-1">Agent Mode</h4>
                <p class="text-xs text-blue-700 dark:text-blue-400">
                  This node was deployed using an agent token. The endpoint URL is automatically detected during registration and cannot be modified.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  node: {
    type: Object,
    default: null
  },
  isEdit: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['close', 'save'])

const saving = ref(false)
const formData = ref({
  name: '',
  location: '',
  url: ''
})

// Initialize form data
if (props.node) {
  formData.value = {
    name: props.node.name || '',
    location: props.node.location || '',
    url: props.node.url || ''
  }
}

// Form validation - only name and location (URL is read-only)
const isFormValid = computed(() => {
  return formData.value.name.trim() && formData.value.location.trim()
})

const handleSubmit = async () => {
  if (!isFormValid.value) return

  saving.value = true
  try {
    emit('save', {
      name: formData.value.name.trim(),
      location: formData.value.location.trim(),
      url: formData.value.url // Keep original URL unchanged
    })
  } finally {
    saving.value = false
  }
}

// Watch for node prop changes
watch(() => props.node, () => {
  if (props.node) {
    formData.value = {
      name: props.node.name || '',
      location: props.node.location || '',
      url: props.node.url || ''
    }
  }
}, { immediate: true })
</script>
