<template>
  <div>
    <v-text-field
      v-model="search"
      label="Search people"
      @input="fetchData"
    />
    <v-data-table
      v-model:items-per-page="itemsPerPage"
      :items="people"
      :headers="headers"
      :items-per-page="15"
    />
  </div>
</template>

<script lang="ts">
import axios from '@/plugins/axios-config'

export default {
  data: () => ({
    people: [],
    headers: [
      { text: "Name", value: "name" },
      { text: "Created", value: "created" },
    ],
    itemsPerPage: 5,
    search: "",
    page: 1,
  }),
  mounted() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      try {
        const result = await axios.get(`/api/people`, { params: { search: this.search, page: this.page } })
        console.log(result)
        if (result) {
          this.people = result.data.results
          console.log(this.people)
        }
      } catch (err) {
        console.error(err)
      }
    },
  },
}
</script>
