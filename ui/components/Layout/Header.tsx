import React from "react";
import { Text, View, StyleSheet } from "react-native";
import { observer } from "mobx-react";

import appStore from "../../stores";

const HeaderComponent = (props: { history: Array<string> }) => {
    const { navigationStore: { componentHeaders } } = appStore;
    const header: string = componentHeaders[props.history.location.pathname];

    return (
        <View style={styles.header}>
            <Text style={styles.headerText}>{header}</Text>
        </View>
    );
};

const styles = StyleSheet.create({
    header: {
      height: "10%",
      width: "100%",
      flexDirection: "row",
      alignItems: "center",
      justifyContent: "center",
      backgroundColor: "#f0f0f0",
      borderBottomColor: "#333",
      borderBottomWidth: 1,
    },
    headerText: {
        fontWeight: "bold",
        fontSize: 20,
        color: "#333",
        letterSpacing: 1,
        paddingTop: 35,
    },
});

export default (observer(HeaderComponent));
