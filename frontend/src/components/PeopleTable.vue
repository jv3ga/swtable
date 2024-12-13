<template>
  <div>
    <v-text-field
      v-model="search"
      label="Search people"
      @input="fetchData"
    />
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="people"
      :items-length="totalItems"
      :loading="loading"
      item-value="name"
      @update:options="fetchData"
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
    itemsPerPage: 15,
    totalItems: 0,
    loading: false,
    search: "",
    page: 1,
  }),
  mounted() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      try {
        const result = await axios.get(`/api/people`, { params: { search: this.search, page: this.page } })
        console.log(result.data.results.length)
        if (result) {
          this.people = result.data.results
          console.log(this.people)
        }
      } catch (err) {
        console.error(err)
      } finally {
        this.loading = false
      }
    },
  },
}
</script>
