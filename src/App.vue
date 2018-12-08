<template>
  <v-app id="inspire">
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-title>Google Domain DDNS Updater</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <ul v-if="errors && errors.length">
              <li v-for="error of errors">
                <v-alert
                  :value="true"
                  type="error"
                >
                  {{error.message}}
                </v-alert>
              </li>
            </ul>
            
            <v-btn color="success">Add</v-btn>
            <v-data-table
              :headers="headers"
              :items="hostnames"
              class="elevation-1"
              hide-actions
            >
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">{{ props.item.domain }}</td>
                <td class="text-xs-left">{{ props.item.username }}</td>
                <td class="text-xs-left">{{ props.item.password }}</td>
                <td>
                  <v-btn color="error">Delete</v-btn>
                  <v-btn color="info">Edit</v-btn>
                </td>
              </template>
            </v-data-table>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="indigo" app>
      <span class="white--text">&nbsp; Visit the <a href="https://github.com/StevenWeathers/google-domains-ddns-updater">Github Repo</a></span>
    </v-footer>
  </v-app>
</template>

<script>
  import axios from 'axios';

  export default {
    data() {
      return {
        headers: [
          {
            text: 'Domain',
            align: 'left',
            value: 'domain'
          },
          {
            text: 'Username',
            sortable: false,
            align: 'left'
          },
          {
            text: 'Password',
            sortable: false,
            align: 'left'
          },
          {
            text: 'Actions',
            sortable: false,
            align: 'left'
          }
        ],
        hostnames: [],
        errors: []
      }
    },
    created() {
      axios.get(`/hostnames`)
      .then(response => {
        // JSON responses are automatically parsed.
        this.hostnames = response.data.hostnames
      })
      .catch(e => {
        this.errors.push(e)
      })
    }
  }
</script>