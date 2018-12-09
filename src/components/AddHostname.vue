<template>
  <v-form ref="form" v-model="valid" lazy-validation>
    <v-text-field v-model="hostname.domain" label="Domain" :rules="domainRules" required></v-text-field>
    <v-text-field v-model="hostname.username" label="Username" :rules="usernameRules" required></v-text-field>
    <v-text-field v-model="hostname.password" label="Password" :rules="passwordRules" required></v-text-field>

    <v-btn color="success" :disabled="!valid" @click="submit">submit</v-btn>
    <v-btn @click="clear">clear</v-btn>
    <v-btn color="error" @click="$emit('closeAddHostname')">cancel</v-btn>
  </v-form>
</template>

<script>
import axios from "axios";

export default {
  name: "AddHostname",
  props: {},
  data() {
    return {
      valid: true,
      domainRules: [
        v => !!v || 'Domain is required',
        v => /^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/.test(v) || 'Domain must be valid'
      ],
      usernameRules: [
        v => !!v || 'Username is required'
      ],
      passwordRules: [
        v => !!v || 'Password is required'
      ],
      hostname: {
        domain: "",
        username: "",
        password: ""
      },
      errors: []
    };
  },
  methods: {
      submit () {
        if (this.$refs.form.validate()) {
          axios
            .post('/hostnames', this.hostname)
            .then(response => {
                this.$emit('closeAddHostname');
            })
            .catch(e => {
                this.errors.push(e);
            })
        }
      },
      clear () {
        this.$refs.form.reset()
      }
  }
};
</script>