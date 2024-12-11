<template>
  <div>
    <v-text-field
      v-model="search"
      label="Search people"
      @input="fetchData"
    />
    <v-data-table
      v-model:page="page"
      :items="people"
      :headers="headers"
      :items-per-page="15"
    />
  </div>
</template>

<script lang="ts">
import axios from "axios"

export default {
  data: () => ({
    people: [],
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
        .get(`/api/people`, {
          params: { search: this.search, page: this.page },
        })
        .then((res) => (this.people = res.data.results))
        .catch((err) => console.error(err))
    },
  },
}
</script>
