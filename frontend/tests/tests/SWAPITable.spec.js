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
      { name: 'Han Solo', created: '2024-01-03T00:00:00Z' },
      { name: 'Darth Vader', created: '2024-01-04T00:00:00Z' },
      { name: 'Obi-Wan Kenobi', created: '2024-01-05T00:00:00Z' },
      { name: 'Yoda', created: '2024-01-06T00:00:00Z' },
      { name: 'Palpatine', created: '2024-01-07T00:00:00Z' },
      { name: 'Lando Calrissian', created: '2024-01-08T00:00:00Z' },
      { name: 'Chewbacca', created: '2024-01-09T00:00:00Z' },
      { name: 'Padmé Amidala', created: '2024-01-10T00:00:00Z' },
      { name: 'Anakin Skywalker', created: '2024-01-11T00:00:00Z' },
      { name: 'R2-D2', created: '2024-01-12T00:00:00Z' },
      { name: 'C-3PO', created: '2024-01-13T00:00:00Z' },
      { name: 'Mace Windu', created: '2024-01-14T00:00:00Z' },
      { name: 'Qui-Gon Jinn', created: '2024-01-15T00:00:00Z' },
      { name: 'Ahsoka Tano', created: '2024-01-16T00:00:00Z' },
      { name: 'Boba Fett', created: '2024-01-17T00:00:00Z' },
      { name: 'Jango Fett', created: '2024-01-18T00:00:00Z' },
      { name: 'Rey', created: '2024-01-19T00:00:00Z' },
      { name: 'Finn', created: '2024-01-20T00:00:00Z' },
      { name: 'Poe Dameron', created: '2024-01-21T00:00:00Z' },
      { name: 'Kylo Ren', created: '2024-01-22T00:00:00Z' },
      { name: 'General Hux', created: '2024-01-23T00:00:00Z' },
      { name: 'Rose Tico', created: '2024-01-24T00:00:00Z' },
      { name: 'Admiral Ackbar', created: '2024-01-25T00:00:00Z' },
      { name: 'ZZZZZ DO NOT SHOW', created: '2024-01-25T00:00:00Z' },
    ],
    count: 25,
  }
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
      // Este no debería estar porque no entra en lo de mostrar 15 elementos solamente
      expect(screen.getByText('ZZZZZ DO NOT SHOW')).toBeInTheDocument()
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
