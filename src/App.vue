<template>
  <v-app id="inspire">
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-title>Google Domains DDNS Updater</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat @click="addHostname" v-show="!showHostnameForm">Add Hostname</v-btn>
        <v-btn flat @click="triggerJob">Trigger Job</v-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout justify-center align-center>
          <v-flex>
            <HostnameForm v-if="showHostnameForm" v-on:closeAddHostname="closeHostnameForm" :is-edit="isHostnameEdit" :edit-hostname="selectedHostname" />
            <HostnamesList v-else v-on:editHostname="editHostname"/>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="indigo" app>
      <span class="white--text">&nbsp; Visit the
        <a href="https://github.com/StevenWeathers/google-domains-ddns-updater">Github Repo</a>
      </span>
    </v-footer>
  </v-app>
</template>

<script>
import axios from "axios"

import HostnamesList from "./components/HostnamesList.vue"
import HostnameForm from "./components/HostnameForm.vue"

export default {
  name: "app",
  data() {
    return {
      showHostnameForm: false,
      isHostnameEdit: false,
      selectedHostname: {},
    }
  },
  components: {
    HostnamesList,
    HostnameForm
  },
   methods: {
    closeHostnameForm() {
      this.isHostnameEdit = false
      this.selectedHostname = {}
      this.showHostnameForm = false
    },
    addHostname() {
      this.isHostnameEdit = false
      this.selectedHostname = {}
      this.showHostnameForm = true
    },
    editHostname(hostname) {
      this.isHostnameEdit = true
      this.selectedHostname = hostname
      this.showHostnameForm = true
    },
    triggerJob() {
      axios
        .get('/triggerUpdate')
    }
  }
};
</script>