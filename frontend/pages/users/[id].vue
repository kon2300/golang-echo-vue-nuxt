<script setup lang="ts">
const searchText = ref('')
const myData = ref([])

const searchForStuff = async () => {
  const { data: fetchData } = await useFetch(`/api/search?search=${searchText.value}`)
  console.log(fetchData.value)
  myData.value = fetchData.value
}
</script>
<template>
  <div>
    <div>
      {{ $route.params }}
    </div>
    <div>
      <form class="form m-1" @submit.prevent="searchForStuff">
        <input class="pl-2 border-2" type="text" v-model="searchText" />
        <button class="border-2 shadow-md bg-blue-100 hover:bg-blue-300">Search</button>
      </form>
      <div v-for="show in myData" :key="show.id">
        <img :src="show.show?.image?.medium" alt="" />
        <p>{{ show.show.name }}</p>
      </div>
    </div>
  </div>
</template>
