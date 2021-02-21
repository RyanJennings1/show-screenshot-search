<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <v-container>
      <v-row no-gutters>
        <v-col cols="12" sm="6" md="6">
          <v-card class="pa-2" outlined tile>
            <MainImage />
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="6">
          <v-card class="pa-2" outlined tile>
            <ImagesCarousel />
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import MainImage from './components/MainImage.vue'
import ImagesCarousel from './components/ImagesCarousel.vue'

export default {
  name: 'App',
  components: {
    MainImage,
    ImagesCarousel
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
