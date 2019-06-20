<template>
  <HelloWorld />
</template>

<script>
  import HelloWorld from '../components/HelloWorld'

  export default {
    components: {
      HelloWorld
    },
    created() {
      this.$router.push({ path: '/user_flights/'});
    }
  }
</script>
