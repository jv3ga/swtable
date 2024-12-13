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
      :items-per-page-options="itemsPerPageOptions"
      @update:options="debouncedFetch"
      @update:sort-by="sortByUpdated"
      @update:page="pageUpdate"
    >
      <template
        #item.created="{ item }"
      >
        <span>{{ $filters.shortDateTime(item.created) }}</span>
      </template>
    </v-data-table-server>
  </div>
</template>

<script lang="ts">
import axios from '@/plugins/axios-config'
import { debounce } from 'lodash'

export default {
  data: () => ({
    people: [],
    headers: [
      { title: "Name", value: "name", aling: "end", sortable: true },
      { title: "Created", value: "created", aling: "end", sortable: true },
    ],
    itemsPerPage: 15,
    totalItems: 0,
    loading: false,
    search: "",
    page: 1,
    sortBy: "name",
    order: "desc",
    itemsPerPageOptions: [
      {value: 15, title: '15'},
    ],
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
    async pageUpdate (value: number) {
      this.page = value
    },
    async fetchData() {
      this.loading = true
      try {
        const result = await axios.get(`/api/people`,
        {
          params: {
            search: this.search,
            page: this.page,
            sortBy: this.sortBy,
            order: this.order,
          }
        })
        if (result) {
          this.people = result.data.results
          this.totalItems = result.data.count
        }
      } catch (err) {
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    async sortByUpdated (value: { key: string; order: string; }[]) {
      const valueSelected = value[0]
      this.sortBy = valueSelected.key
      this.order = valueSelected.order
    }
  },
}
</script>
