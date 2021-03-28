<template>
    <div class="container is-desktop">
        <div class="number-options-board tile is-ancestor">
            <number-option-box v-for="(item, i) in numbers" :key="i" :number="item">
            </number-option-box>
        </div>
    </div>
</template>

<script>
    import NumberOptionBox from "./NumberOptionBox";

    export default {
        name: "SelectNumberBoard",
        components: {
            "number-option-box": NumberOptionBox
        },
        data () {
            return {
                numbers: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
            }
        }
    }
</script>

<style scoped>
    .number-options-board {
        display: flex;
        flex-wrap: wrap;
    }
</style>
