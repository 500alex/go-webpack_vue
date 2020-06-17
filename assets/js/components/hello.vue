<template>
<div>
    <div>Hello from component!</div>
    <div>{{ incomponent }}</div>
</div>

</template>

<script>
export default {
    data() {
        return {
            incomponent: 'text in component'
        }
    },
    created() {
        console.log("hello in hello create")
    }
}
</script>

<style lang="scss">
body {
    background: #cecece;
}
</style>

