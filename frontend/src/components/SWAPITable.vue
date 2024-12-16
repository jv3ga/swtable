<template>
  <div
    v-if="errorMessage"
  >
    <v-alert
      type="error"
      :text="errorMessage"
    />
  </div>
  <div
    v-else
  >
    <v-text-field
      v-model="search"
      label="Search items"
      @input="debouncedFetch"
    />
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="items"
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
  props: {
    apiUrl: {
      type: String,
      required: true,
    }
  },
  data: () => ({
    items: [],
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
    errorMessage: '',
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
        const result = await axios.get(this.apiUrl,
        {
          params: {
            search: this.search,
            page: this.page,
            sortBy: this.sortBy,
            order: this.order,
          }
        })
        if (result) {
          this.items = result.data.results
          this.totalItems = result.data.count
        }
      } catch (error: debounce) {
        if (error.response) {
          // El servidor respondió con un código de estado que no está en el rango 2xx
          console.error("Error status:", error.response.status)
          console.error("Error data:", error.response.data)
          console.error("Headers:", error.response.headers)
          this.errorMessage = error.response.data
        } else if (error.request) {
          // La solicitud fue enviada pero no hubo respuesta
          this.errorMessage = `Error request: ${error.request}`
        } else {
          // Algo ocurrió al configurar la solicitud
          this.errorMessage = `Axios error: ${error.message}`
        }
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
