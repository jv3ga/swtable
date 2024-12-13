import filters from '@/utils/filters'

export default {
  install(app: any) {
    // Registrar filtros como propiedades globales
    app.config.globalProperties.$filters = {
      shortDateTime: filters.shortDateTime,
    }
  }
}
