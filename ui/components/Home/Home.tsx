import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

const styles = StyleSheet.create({
    component: {
        height: "80%",
    },
});

const HomeComponent = (prop: { history: Array<string> }) => (
    <View style={styles.component}>
        <Text>This is the home page.</Text>
        <Button title="Change Page" onPress={() => prop.history.push("/profile")} />
    </View>
);

export default HomeComponent;
