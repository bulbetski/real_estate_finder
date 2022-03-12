<script>
export default {
    data() {
        return {
            count: 0
        }
    },
    methods: {
        increment() {
            this.count++
        }
    }
}
</script>

<template>
    <button @click="this.increment()">count is: {{ count }}</button>
</template>