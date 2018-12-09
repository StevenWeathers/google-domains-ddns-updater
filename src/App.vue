<template>
  <v-app id="inspire">
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-title>Google Domain DDNS Updater</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat @click="addHostname" v-show="!showAddHostname">Add Hostname</v-btn>
        <v-btn flat @click="triggerJob">Trigger Job</v-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout justify-center align-center>
          <v-flex>
            <AddHostname v-if="showAddHostname" v-on:closeAddHostname="showAddHostname = false" />
            <HostnamesList v-else/>
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
import axios from "axios";

import HostnamesList from "./components/HostnamesList.vue";
import AddHostname from "./components/AddHostname.vue";

export default {
  name: "app",
  data() {
    return {
      hostname: {},
      showAddHostname: false,
    }
  },
  components: {
    HostnamesList,
    AddHostname
  },
   methods: {
    addHostname() {
      this.showAddHostname = true;
    },
    editHostname(hostname) {
      this.hostname = hostname;
    },
    triggerJob() {
      axios
        .get('/triggerUpdate')
    }
  }
};
</script>