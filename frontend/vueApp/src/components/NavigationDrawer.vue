<template>
    <v-navigation-drawer
            v-model="drawer"
            :mini-variant.sync="mini"
            permanent
            app
    >
        <v-toolbar flat class="transparent">
            <v-list>
                <v-list-tile avatar>
                    <v-list-tile-avatar>
                        <v-btn
                                icon
                                @click.stop="sizeClick"
                        >
                            <v-icon>menu</v-icon>
                        </v-btn>
                    </v-list-tile-avatar>

                    <v-list-tile-content>
                        <v-list-tile-title>{{title}}</v-list-tile-title>
                    </v-list-tile-content>

                </v-list-tile>
            </v-list>
        </v-toolbar>

        <v-list dense>
            <v-divider></v-divider>

            <v-list-tile
                    v-for="item in items"
                    :key="item.title"
                    @click="visit(item.route)"
            >
                <v-list-tile-action>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-tile-action>

                <v-list-tile-content>
                    <v-list-tile-title>{{ item.title }}</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
            <v-list-tile @click="logout()">
                <v-list-tile-action>
                    <v-icon>power_settings_new</v-icon>
                </v-list-tile-action>

                <v-list-tile-content>
                    <v-list-tile-title>Logout</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
    export default {
        name: "NavigationDrawer",
        props: ["title", "items"],
        data () {
            return {
                drawer: true,
                mini: true,
            }
        },
        methods: {
            logout(){
                this.$cookies.remove("token");
                this.$router.replace("/admin_login")
            },
            visit(path) {
                this.$router.replace(path);
            },
            sizeClick() {
                this.mini = !this.mini;
                this.$emit('sizeClick', this.mini);
            }
        }
    }
</script>

<style scoped>

</style>