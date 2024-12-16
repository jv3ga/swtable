import { render, fireEvent, screen, waitFor } from '@testing-library/vue';
import SWAPITable from '@/components/SWAPITable.vue';
import { vi } from 'vitest';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import axiosInstance from '@/plugins/axios-config';

// Mock necesario para ResizeObserver en Vuetify
global.ResizeObserver = require('resize-observer-polyfill')

// Mockea la instancia de Axios
vi.mock('@/plugins/axios-config', () => ({
  default: {
    get: vi.fn(), // Mockea `get`
  },
}))

describe('SWAPITable Component Tests', () => {
  const apiUrl = "/api/people";

  const mockData = {
    results: [
      { name: 'Luke Skywalker', created: '2024-01-01T00:00:00Z' },
      { name: 'Leia Organa', created: '2024-01-02T00:00:00Z' },
    ],
    count: 2,
  };
  const vuetify = createVuetify({
    components,
    directives,
  })

  beforeEach(() => {
    axiosInstance.get.mockReset()
    axiosInstance.get.mockResolvedValue({ data: mockData })
  })

  const defaultComponentRenderOptions = {
    props: { apiUrl },
    global: {
      plugins: [vuetify],
      mocks: {
        $filters: {
          shortDateTime: (date) => `Formatted: ${date}`,
        },
      },
    },
  }

  it('displays fetched data correctly', async () => {
    // Renderiza el componente con las props necesarias
    render(SWAPITable, {
      props: { apiUrl },
      global: {
        plugins: [vuetify],
        mocks: {
          $filters: {
            shortDateTime: (date) => `Formatted: ${date}`,
          },
        },
      },
    })

    // Verifica que `axiosInstance.get` haya sido llamado correctamente
    await waitFor(() => {
      expect(axiosInstance.get).toHaveBeenCalledWith(apiUrl, {
        params: {
          search: "",
          page: 1,
          sortBy: "name",
          order: "desc",
        },
      })
    })

    await waitFor(() => {
      expect(screen.getByText('Luke Skywalker')).toBeInTheDocument()
      expect(screen.getByText('Leia Organa')).toBeInTheDocument()
    })
  })

  it('displays an error message on API failure', async () => {
    axiosInstance.get.mockRejectedValueOnce({ response: { data: 'API Error' } })

    render(SWAPITable, defaultComponentRenderOptions)

    await waitFor(() => {
      expect(screen.getByText('API Error')).toBeInTheDocument()
    })
  })

  it('filters items based on search input', async () => {
    axiosInstance.get.mockResolvedValue(mockData)

    render(SWAPITable, defaultComponentRenderOptions)

    const searchField = screen.getByLabelText('Search items')

    await fireEvent.update(searchField, 'Luke')
    await waitFor(() => {
      expect(axiosInstance.get).toHaveBeenCalledWith(apiUrl, expect.objectContaining({
        params: expect.objectContaining({ search: 'Luke' }),
      }))
    })
  })
  // TODO: check how the sort and next page buttons are rendered
  it('paginates correctly', async () => {
    axiosInstance.get.mockResolvedValueOnce(mockData)

    render(SWAPITable, defaultComponentRenderOptions)

    await fireEvent.click(screen.getByRole('button', { name: /next/i }))

    await waitFor(() => {
      expect(axiosInstance.get).toHaveBeenCalledWith(apiUrl, expect.objectContaining({
        params: expect.objectContaining({ page: 2 }),
      }))
    })
  })
  it('supports sorting by name and created fields', async () => {
    axiosInstance.get.mockResolvedValue(mockData)
    render(SWAPITable, defaultComponentRenderOptions)
    const sortByNameButton = screen.getByText('Name');
    const sortByCreatedButton = screen.getByText('Created');

    await fireEvent.click(sortByNameButton);
    await waitFor(() => {
      expect(axiosInstance.get).toHaveBeenCalledWith(apiUrl, expect.objectContaining({
        params: expect.objectContaining({ sortBy: 'name', order: 'asc' }),
      }));
    });

    await fireEvent.click(sortByCreatedButton);
    await waitFor(() => {
      expect(axiosInstance.get).toHaveBeenCalledWith(apiUrl, expect.objectContaining({
        params: expect.objectContaining({ sortBy: 'created', order: 'asc' }),
      }));
    });
  });

})
