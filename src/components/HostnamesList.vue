<template>
  <div>
    <template v-if="errors && errors.length">
      <v-alert :value="true" type="error" v-for="error of errors" :key="error.message">{{message}}</v-alert>
    </template>
    <v-data-table :headers="headers" :items="hostnames" class="elevation-1" hide-actions>
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{ props.item.domain }}</td>
        <td class="text-xs-left">{{ props.item.username }}</td>
        <td class="text-xs-left">{{ props.item.password }}</td>
        <td>
          <v-btn color="info" @click="$emit('editHostname', props.item)">Edit</v-btn>
          <v-btn color="error" @click="handleDelete(props.item.domain)">Delete</v-btn>
        </td>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "HostnamesList",
  props: {},
  data() {
    return {
      headers: [
        {
          text: "Domain",
          align: "left",
          value: "domain"
        },
        {
          text: "Username",
          sortable: false,
          align: "left"
        },
        {
          text: "Password",
          sortable: false,
          align: "left"
        },
        {
          text: "Actions",
          sortable: false,
          align: "left"
        }
      ],
      hostnames: [],
      errors: []
    };
  },
  created() {
    axios
      .get(`/hostnames`)
      .then(response => {
        this.hostnames = response.data.hostnames
      })
      .catch(e => {
        this.errors.push(e)
      });
  },
  methods: {
    handleDelete(domain) {
      axios
        .delete(`/hostnames/${domain}`)
        .then(response => {
          const hostnameIndex = this.hostnames.findIndex(hostname => hostname.domain === domain)
          this.hostnames.splice(hostnameIndex, 1)
        })
        .catch(e => {
          this.errors.push(e)
        })
    },
  }
};
</script>