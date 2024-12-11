<template>
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
          <PeopleTable />
        </v-tabs-window-item>
        <v-tabs-window-item value="planets">
          <PlanetsTable />
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import PeopleTable from './components/PeopleTable.vue'
import PlanetsTable from './components/PlanetsTable.vue'

export default {
  components: { PeopleTable, PlanetsTable },
  data: () => ({
      tab: null,
    }),
}
</script>
