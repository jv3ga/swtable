<template>
  <v-container>
    <v-card>
      <v-tabs
        v-model="tab"
        bg-color="primary"
      >
        <v-tab value="people">
          People
        </v-tab>
        <v-tab value="planets">
          Planets
        </v-tab>
      </v-tabs>
      <v-card-text>
        <v-tabs-window v-model="tab">
          <v-tabs-window-item value="people">
            <SWAPITable :api-url="API_URLS.people" />
          </v-tabs-window-item>
          <v-tabs-window-item value="planets">
            <SWAPITable :api-url="API_URLS.planets" />
          </v-tabs-window-item>
        </v-tabs-window>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import SWAPITable from "./components/SWAPITable.vue"

const API_URLS = {
  people: "/api/people",
  planets: "/api/planets",
}

export default {
  components: { SWAPITable },
  data: () => ({
    tab: "people",
  }),
  computed: {
    API_URLS() {
      return API_URLS
    },
  },
}
</script>
