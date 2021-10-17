import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

const styles = StyleSheet.create({
    component: {
        height: "80%",
    },
});

const SettingsComponent = (prop: { history: Array<string> }) => (
    <View style={styles.component}>
        <Text>This is the settings component</Text>
        <Text>More text</Text>
        <Button title="Button" onPress={() => 5} />
    </View>
);

export default SettingsComponent;
