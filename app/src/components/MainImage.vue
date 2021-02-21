<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <img class="mainImage" :src="getImageUrl('')"/>
    <p>
      Main image here
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
      count: 0
    }
  },
  methods: {
		getImageUrl(im) {
			var images = require.context('../../../frames/', false, /\.png$/)
			console.log(im);
			//return images('./' + im + '.png')
			return images('./test0001.png')
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
.mainImage {
  width: 60%
}
</style>
