import QtQuick
import QtQuick.Controls
import qs.Commons

// A wider, easier-to-grab replacement for the platform-default ScrollBar,
// which renders as a near-invisible sliver on most Omarchy themes.
ScrollBar {
    id: root

    policy: ScrollBar.AsNeeded
    implicitWidth: Style.space(12)
    implicitHeight: Style.space(12)

    contentItem: Rectangle {
        implicitWidth: Style.space(8)
        implicitHeight: Style.space(8)
        radius: width / 2
        color: root.pressed
            ? Color.accent
            : Util.alpha(Color.accent, root.hovered ? 0.85 : 0.55)
    }

    background: Rectangle {
        implicitWidth: Style.space(12)
        implicitHeight: Style.space(12)
        radius: width / 2
        color: Util.alpha(Color.accent, 0.08)
    }
}
