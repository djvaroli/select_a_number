<template>
    <section>
        <b-steps type="is-success">
            <b-step-item label="Pick a Number" icon="tag">
                <div class="flex-center-wrapper">
                    <div class="step-text-prompt box">
                        Here's the deal. All you have to do for this task is pick a number. You can pick any number
                        starting at 1 and ending at 10. Simple!<br><br>
                        However, you have to pick a number before you contiue. So, here's your chance. Pick a number, write
                        it down and then feel free to hit that nice looking continue button!
                    </div>
                </div>
            </b-step-item>
            <b-step-item label="Some Info" icon="account">
                <div class="flex-center-wrapper">
                    <div class="width-300">
                        <b-field label="Age">
                            <b-input v-model="age"
                                     placeholder="Your age"
                                     type="number"
                                     icon-pack="fas"
                                     icon="user-alt"
                            >

                            </b-input>
                        </b-field>

                        <b-field label="Nationality">
                            <b-input v-model="nationality"
                                     placeholder="city you call home..."
                                     type="text"
                            >
                            </b-input>
                        </b-field>

                        <b-field label="Gender">
                            <b-input v-model="gender"
                                     placeholder="self-explanatory"
                                     type="text"
                            >
                            </b-input>
                        </b-field>

                        <b-field label="Dominant Hand">
                            <b-input v-model="dominantHand"
                                     placeholder="which hand you use most often..."
                                     type="text"
                            >
                            </b-input>
                        </b-field>
                    </div>
                </div>
            </b-step-item>
            <b-step-item label="Submit Your Choice" icon="create">
                <div class="flex-center-wrapper">
                    <div class="step-text-prompt">
                        You made it! Time to click on the box that corresponds to the number you chose before!
                    </div>
                    <div class="container is-desktop">
                        <div class="number-options-board tile is-ancestor">
                            <number-option-box v-for="(item, i) in numbers" :key="i" :number="item">
                            </number-option-box>
                        </div>
                    </div>
                </div>
            </b-step-item>

            <template #navigation="{previous, next}">
                <b-button class="step-navigation-button"
                        v-if="showBackButton"
                        outlined
                        type="is-danger"
                        icon-pack="fas"
                        icon-right="backward"
                        :disabled="previous.disabled"
                        @click.prevent="previous.action">
                    Go Back
                </b-button>
                <b-button class="step-navigation-button"
                        outlined
                        type="is-primary"
                        icon-pack="fas"
                        icon-right="forward"
                        :disabled="next.disabled"
                        @click="showBackButton=true"
                        @click.prevent="next.action">
                    Continue
                </b-button>
            </template>
        </b-steps>
    </section>


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
                numbers: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                showBackButton: false
            }
        }
    }
</script>

<style scoped lang="scss">

    .width-300 {
        max-width: 300px;
    }
    .number-options-board {
        display: flex;
        flex-wrap: wrap;
    }

    .flex-center-wrapper{
        display: flex;
        justify-content: center;
        text-align: center;
        .step-text-prompt {
            max-width: 400px;
        }
    }

    .step-navigation-button {
        margin: 10px;
    }



</style>
