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
            <SWAPITable
              api-url="/api/people"
            />
          </v-tabs-window-item>
          <v-tabs-window-item value="planets">
            <SWAPITable
              api-url="/api/planets"
            />
          </v-tabs-window-item>
        </v-tabs-window>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import SWAPITable from './components/SWAPITable.vue';

export default {
  components: { SWAPITable },
  data: () => ({
      tab: null,
    }),
}
</script>
