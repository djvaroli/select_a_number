<template>
    <section>
        <b-steps type="is-primary" mobile-mode="minimalist">
            <b-step-item label="Pick a Number" step="1">
                <div class="flex-content-center-wrapper text-left">
                    <div class="step-text-prompt box">
                        Here's the deal. All you have to do for this task is pick a number. You can pick any number
                        starting at 1 and ending at 10. Simple!<br><br>
                        However, you have to pick a number before you contiue. So, here's your chance. Pick a number, write
                        it down and then feel free to hit that nice looking continue button!
                    </div>
                </div>
            </b-step-item>
            <b-step-item label="Some Info" step="2">
                <div class="flex-content-center-wrapper">
                    <div class="text-left box stretch-max-500">
                        <b-field label="Age">
                            <b-input v-model="age"
                                     placeholder="Your age"
                                     type="number"
                                     icon-pack="fas"
                                     icon="user-alt"
                            >

                            </b-input>
                        </b-field>

                        <b-field label="Country" required>
                            <b-input v-model="nationality"
                                     placeholder="place you call home..."
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
            <b-step-item label="Submit Your Choice" step="3">
                <div class="flex-content-center-wrapper">
                    <div class="text-left box stretch-max-500">
                        <b-field label="Your number">
                            <b-input v-model="numberChoice"
                                     placeholder="That number you wrote down..."
                                     type="number"
                                     icon-pack="fas"
                                     icon="user-alt"
                            >
                            </b-input>
                        </b-field>
                    </div>
                </div>
            </b-step-item>
            <b-step-item label="Done!" step="4" type="is-success">
                <div class="flex-content-center-wrapper">
                    <b-button type="is-success" size="is-large" class="button-box" outlined>Submit</b-button>
                </div>
            </b-step-item>

<!--            <template #navigation="{previous, next}">-->
<!--                <b-button class="step-navigation-button"-->
<!--                        v-if="showBackButton"-->
<!--                        outlined-->
<!--                        type="is-danger"-->
<!--                        icon-pack="fas"-->
<!--                        icon-right="backward"-->
<!--                        :disabled="previous.disabled"-->
<!--                        @click.prevent="previous.action">-->
<!--                    Go Back-->
<!--                </b-button>-->
<!--                <b-button class="step-navigation-button"-->
<!--                        outlined-->
<!--                        type="is-primary"-->
<!--                        icon-pack="fas"-->
<!--                        icon-right="forward"-->
<!--                        :disabled="next.disabled"-->
<!--                        @click="showBackButton=true"-->
<!--                        @click.prevent="next.action">-->
<!--                    Continue-->
<!--                </b-button>-->
<!--            </template>-->
        </b-steps>
    </section>


</template>

<script>
    // import NumberOptionBox from "./NumberOptionBox";

    export default {
        name: "SelectNumberBoard",
        components: {
            // "number-option-box": NumberOptionBox
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

    .number-options-board {
        display: flex;
        flex-wrap: wrap;
    }

    .text-left {
        text-align: left;
    }

    .text-right {
        text-align: right
    }

    .margin-3rem {
        margin: 3rem;
    }

    .max-width-500 {
        max-width: 500px;
    }

    .stretch-max-500 {
        width: 100%;
        max-width: 500px;
    }

    .flex-content-center-wrapper{
        display: flex;
        margin: 3rem 0;
        justify-content: center;
        .step-text-prompt {
            width: 100%;
            max-width: 500px;
        }
    }

    .step-navigation-button {
        margin: 10px;
    }

    .button-box {
        box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    }





</style>
