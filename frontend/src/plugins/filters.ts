import filters from '@/utils/filters'

export default {
  install(app) {
    // Registrar filtros como propiedades globales
    app.config.globalProperties.$filters = {
      shortDateTime: filters.shortDateTime,
    }
  }
}
