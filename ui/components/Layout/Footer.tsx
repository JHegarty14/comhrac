import React from "react";
import { observer, inject } from "mobx-react";
import { View, Text, StyleSheet, GestureResponderEvent } from "react-native";
import { Icon } from "react-native-elements";

const FooterComponent = (props: { history: Array<string> }) => {

    const handlePress = (e: GestureResponderEvent, path: string): void => {
        console.log("PATH", path);
        props.history.push(path);
    };

    return (
        <View style={styles.footer}>
            <View style={styles.iconRow}>
                <Icon
                    type="material-community"
                    name="calendar"
                    onPress={(e) => handlePress(e, "/progress")}
                    color={props.history.location.pathname === "/progress" ? "#ed93ea" : ""}
                />
                <Icon
                    type="material-community"
                    name="av-timer"
                    onPress={(e) => handlePress(e, "/timer")}
                    color={props.history.location.pathname === "/timer" ? "#ed93ea" : ""}
                />
                <Icon
                    type="material-community"
                    name="dumbbell"
                    onPress={(e) => handlePress(e, "/workouts")}
                    color={props.history.location.pathname === "/workouts" ? "#ed93ea" : ""}
                />
                <Icon
                    type="material-community"
                    name="pencil-plus-outline"
                    onPress={(e) => handlePress(e, "/workouts/create-workout")}
                    color={props.history.location.pathname === "/workouts/create-workout" ? "#ed93ea" : ""}
                />
                <Icon
                    type="material-community"
                    name="account-outline"
                    onPress={(e) => handlePress(e, "/profile")}
                    color={props.history.location.pathname === "/profile" ? "#ed93ea" : ""}
                />
            </View>
        </View>
    );
};

const styles = StyleSheet.create({
    footer: {
      height: "10%",
      width: "100%",
      backgroundColor: "#f0f0f0",
      borderTopColor: "#333",
      borderTopWidth: 1,
    },
    test: {

    },
    iconRow: {
        height: "75%",
        flexDirection: "row",
        flexWrap: "nowrap",
        alignItems: "center",
        justifyContent: "space-around",
        // paddingBottom: "0",
    },
    footerText: {
        fontWeight: "bold",
        fontSize: 20,
        color: "#333",
        letterSpacing: 1,
    },
});

export default (observer(FooterComponent));
