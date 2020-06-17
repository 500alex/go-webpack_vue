<template>
    <div id="app">
        <div>Hello in app vue</div>
        <hello/>
        <div>{{ message}}</div>

    </div>
</template>



<script>
import Hello from './js/components/hello'
export default {
    data() {
        return {
            message: "HI!"
        }
    },
    created() {
        console.log('hello from app create')
    },
    components: {
        'hello': Hello
    }
}
</script>

