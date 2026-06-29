<script setup lang="ts">
import { Events } from "@wailsio/runtime";

const gap = 16;
const scrollArea = useTemplateRef("scrollArea");
const { width } = useElementSize(() => scrollArea.value?.$el);

const lanes = computed(() =>
  Math.max(1, Math.min(4, Math.floor(width.value / 400))),
);
const laneWidth = computed(
  () => (width.value - (lanes.value - 1) * gap) / lanes.value,
);
const estimateSize = computed(() => laneWidth.value * (480 / 640));

const todoGroupList = ref<TodoGroup[]>([]);
let unsubscribe: () => void;

async function getAllTodoG() {
  const todoG = await TodoService.GetAllTodoGroup();
  if (todoG) {
    todoGroupList.value = todoG;
  }
}

watch(width, (l) => {
  console.log(l, lanes.value);
});

onMounted(() => {
  getAllTodoG();

  unsubscribe = Events.On("todoService:changes", getAllTodoG);
});

onUnmounted(() => {
  unsubscribe();
});
</script>

<template>
  <div class="flex flex-col w-full h-full items-center">
    <UScrollArea
      ref="scrollArea"
      class="w-full h-full p-2"
      :items="todoGroupList"
      v-slot="{ item }"
      :virtualize="{
        gap,
        lanes,
        estimateSize,
      }"
    >
      <TodoGroupCard :todo-group="item" />
    </UScrollArea>
  </div>
</template>
