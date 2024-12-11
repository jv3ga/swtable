<template>
  <div>
    <v-text-field
      v-model="search"
      label="Search planets"
      @input="fetchData"
    />
    <v-data-table
      v-model:page="page"
      :items="planets"
      :headers="headers"
      :items-per-page="15"
    />
  </div>
</template>

<script lang="ts">
import axios from "axios"

export default {
  data: () => ({
    planets: [],
    headers: [
      { text: "Name", value: "name" },
      { text: "Created", value: "created" },
    ],
    search: "",
    page: 1,
  }),
  mounted() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      axios
        .get(`/api/planets`, {
          params: { search: this.search, page: this.page },
        })
        .then((res) => (this.planets = res.data.results))
        .catch((err) => console.error(err))
    },
  },
}
</script>
