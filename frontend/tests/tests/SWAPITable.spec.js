import { mount } from '@vue/test-utils'
import { expect, test } from 'vitest'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import SWAPITable from '@/components/SWAPITable.vue'

const vuetify = createVuetify({
  components,
  directives,
})

global.ResizeObserver = require('resize-observer-polyfill')

test('displays message', () => {
  const wrapper = mount({
    template: '<SWAPITable></SWAPITable>'
  }, {
    props: {},
    global: {
      components: {
        SWAPITable,
      },
      props: {
        apiUrl:"/api/people"
      },
      plugins: [vuetify],
    }
  })

  // Assert the rendered text of the component
  expect(wrapper.text()).toContain('Search items')
})
