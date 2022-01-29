import { View, FlatList, StyleSheet, Text, StatusBar } from "react-native";
import React from "react";
import { useQuery, gql } from "@apollo/client";

const User = ({ name }) => (
  <View style={styles.item}>
    <Text style={styles.title}>{name}</Text>
  </View>
);

export default function UserScreen() {
  const { loading, error, data } = useQuery(
    gql`
      query {
        allUsers {
          name
          id
        }
      }
    `
  );

  if (loading) return <Text>Loading...</Text>;
  if (error) {
    console.log(error);
    return <Text>Error loading users</Text>;
  }

  // TODO generate gql types and apply to data data to ensure safety
  data.allUsers.append({ name: "+ Add User", id: -1 });

  return (
    <FlatList
      ListHeaderComponent={<Text>Users</Text>}
      data={data.allUsers}
      renderItem={({ item }) => <User name={item.name} />}
      keyExtractor={(item) => item.id}
    />
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    marginTop: StatusBar.currentHeight || 0,
  },
  item: {
    backgroundColor: "#f9c2ff",
    padding: 20,
    marginVertical: 8,
    marginHorizontal: 16,
  },
  title: {
    fontSize: 32,
  },
});
