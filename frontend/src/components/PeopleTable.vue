<template>
  <div>
    <v-text-field
      v-model="search"
      label="Search people"
      @input="debouncedFetch"
    />
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="people"
      :items-length="totalItems"
      :loading="loading"
      item-value="name"
      @update:options="debouncedFetch"
    />
  </div>
</template>

<script lang="ts">
import axios from '@/plugins/axios-config'
import { debounce } from 'lodash'

export default {
  data: () => ({
    people: [],
    headers: [
      { title: "Name", value: "name", aling: "end" },
      { title: "Created", value: "created", aling: "end" },
    ],
    itemsPerPage: 15,
    totalItems: 0,
    loading: false,
    search: "",
    page: 1,
    debouncedFetch: () => {},
    seachDelayMs: 500,
  }),
  mounted() {
    this.fetchData()
  },
  created() {
    this.debouncedFetch = debounce(this.fetchData, this.seachDelayMs)
  },
  methods: {
    async fetchData() {
      console.log('fetching...')
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
