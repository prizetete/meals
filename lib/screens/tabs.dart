import 'package:flutter/material.dart';
import 'package:meals/screens/categories.dart';
import 'package:meals/screens/filter.dart';
import 'package:meals/screens/meals.dart';
import 'package:meals/widgets/main_drawer.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:meals/providers/favorites_provider.dart';
import 'package:meals/providers/filters_provider.dart';

const kInitialFilters = {
  Filter.glutenFree: false,
  Filter.lactoseFree: false,
  Filter.vegetarian: false,
  Filter.vegan: false,
};

class TabsScreen extends ConsumerStatefulWidget {
  const TabsScreen({super.key});

  @override
  ConsumerState<TabsScreen> createState() => _TabsScreenState();
}

class _TabsScreenState extends ConsumerState<TabsScreen> {
  var currentIndex = 0;
  // Map<Filter, bool> _selectedFilters = kInitialFilters;

  void selectTab(int selectedIndex) {
    setState(() {
      currentIndex = selectedIndex;
    });
  }

  void _onSelectScreen(String identifier) async {
    Navigator.of(context).pop();
    if (identifier == 'filters') {
      // final result = await Navigator.of(context).push<Map<Filter, bool>>(
      //   MaterialPageRoute(
      //     builder: (context) => FilterScreen(currentFilters: _selectedFilters),
      //   ),
      // );
      // setState(() {
      //   _selectedFilters = result ?? kInitialFilters;
      // });

      await Navigator.of(context).push<Map<Filter, bool>>(
        MaterialPageRoute(builder: (context) => FilterScreen()),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final availableMeals = ref.watch(filteredMealProvider);

    Widget activeWidget = CategoriesScreen(availableMeals: availableMeals);
    var activePageTitle = 'Categories';

    if (currentIndex == 1) {
      final favoriteMeals = ref.watch(favoriteMealsProvider);
      activeWidget = MealsScreen(meals: favoriteMeals);
      activePageTitle = 'Favorites';
    }

    return Scaffold(
      appBar: AppBar(title: Text(activePageTitle)),
      drawer: MainDrawer(onSelectScreen: _onSelectScreen),
      body: activeWidget,
      bottomNavigationBar: BottomNavigationBar(
        onTap: selectTab,
        currentIndex: currentIndex,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.set_meal),
            label: 'Categories',
          ),
          BottomNavigationBarItem(icon: Icon(Icons.star), label: 'Favorites'),
        ],
      ),
    );
  }
}
