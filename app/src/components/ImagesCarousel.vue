<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div>
      <template v-for="pic in getAllImages()">
        <img class="carouselImage"
             :src="getImageUrl(pic)"
             v-bind:alt="pic"
             v-bind:key="pic"/>
      </template>
    </div>
    <p>
      Images here
    </p>
  </div>
</template>

<script>
export default {
  name: 'SelectedImage',
  props: {
    msg: String
  },
  data: function() {
    return {
      count: 0,
    }
  },
  methods: {
		getImageUrl(im) {
      var images = require.context('../../../frames/', false, /\.png$/)
      return images(im)
		},
    getAllImages() {
			var images = require.context('../../../frames/', false, /\.png$/)
			return images.keys();
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.carouselImage {
  width: 25%
}
</style>
