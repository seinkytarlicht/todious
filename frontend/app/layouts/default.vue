<script setup lang="ts">
import { Window, Events } from "@wailsio/runtime";

const isMaximised = ref(false);

function minimize() {
  Window.Minimise();
}

async function toggleMaximize() {
  const isIt = await Window.IsMaximised();

  if (isIt) {
    Window.UnMaximise();
    isMaximised.value = false;
  } else {
    Window.Maximise();
    isMaximised.value = true;
  }
}

function close() {
  Window.Close();
}
</script>

<template>
  <main class="h-screen w-full flex flex-col overflow-hidden">
    <header
      style="--wails-draggable: drag"
      class="w-full flex items-center p-2 justify-between"
    >
      <div>
        <h1 class="text-xl font-semibold ms-3 select-none">
          <span class="text-primary-500">to</span>dious
        </h1>
      </div>

      <div class="w-full max-w-100" style="--wails-draggable: no-drag">
        <CreateTodoGroup />
      </div>

      <div style="--wails-draggable: no-drag">
        <UButton
          color="neutral"
          variant="ghost"
          icon="i-ic-baseline-minimize"
          aria-label="Minimize"
          @click="minimize"
        />
        <UButton
          color="neutral"
          variant="ghost"
          icon="i-ic-baseline-crop-square"
          aria-label="Fullscreen"
          @click="toggleMaximize"
        />

        <UButton
          color="neutral"
          variant="ghost"
          icon="i-ic-round-close"
          aria-label="Close"
          @click="close"
        />
      </div>
    </header>

    <div class="flex-1 min-h-0 overflow-hidden p-4">
      <slot />
    </div>
  </main>
</template>
