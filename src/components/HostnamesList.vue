<template>
  <div>
    <ul v-if="errors && errors.length">
      <li v-for="error of errors" :key="error.message">
        <v-alert :value="true" type="error">{{error.message}}</v-alert>
      </li>
    </ul>
    <v-data-table :headers="headers" :items="hostnames" class="elevation-1" hide-actions>
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{ props.item.domain }}</td>
        <td class="text-xs-left">{{ props.item.username }}</td>
        <td class="text-xs-left">{{ props.item.password }}</td>
        <td>
          <v-btn color="info" @click="handleEdit(props.item.domain)">Edit</v-btn>
          <v-btn color="error" @click="handleDelete(props.item.domain, props.index)">Delete</v-btn>
        </td>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios";

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
        this.hostnames = response.data.hostnames;
      })
      .catch(e => {
        this.errors.push(e);
      });
  },
  methods: {
    handleEdit() {
      // @TODO - hook this up to a form (inline or like Add)
    },
    handleDelete(domain, index) {
      axios
        .delete(`/hostnames/${domain}`)
        .then(response => {
          this.hostnames.splice(index, 1);
        })
        .catch(e => {
          this.errors.push(e);
        })
    },
  }
};
</script>