<template>
  <div class="deployments">
    <h1>This is the deployments page</h1>
    <table class="table-auto w-full bg-white border border-gray-300">
      <thead class="table-fixed bg-gray-800 text-white">
        <tr>
          <th>ID</th>
          <th>Start</th>
          <th>End</th>
          <th>Service</th>
          <th>Environment</th>
          <th>Country</th>
          <th>Source</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in data" :key="item.ID">
          <td>{{item.ID }}</td>
          <td>{{item.Start }}</td>
          <td>{{item.End }}</td>
          <td>{{item.Service }}</td>
          <td>{{item.Environment }}</td>
          <td>{{item.Country }}</td>
          <td>{{item.Source }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      data: [],
    };
  },
  mounted() {
    this.fetchDataFromAPI();
  },
  methods: {
    fetchDataFromAPI() {
      const apiUrl = 'http://localhost:8080/deployments?limit=500'; // Replace with your API endpoint URL

      axios.get(apiUrl)
        .then((response) => {
          this.data = response.data;
        })
        .catch((error) => {
          console.error('Error fetching data:', error);
        });
    },
  },
};
  </script>
