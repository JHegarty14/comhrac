import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

const styles = StyleSheet.create({
    component: {
        height: "80%",
    },
});

const CreateCustomWorkoutComponent = (prop: { history: Array<string> }) => (
    <View style={styles.component}>
        <Text>This is the create custom workout page.</Text>
        <Button title="Change Page" onPress={() => prop.history.push("/profile")} />
    </View>
);

export default CreateCustomWorkoutComponent;
