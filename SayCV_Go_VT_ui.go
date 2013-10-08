// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main

import (
	"github.com/lxn/walk"
)

type dialogUI struct {
	textEdit                *walk.TextEdit
	tabWidget               *walk.TabWidget
	displayTab              *walk.TabPage
	displayAsGroupBox       *walk.GroupBox
	portTab                 *walk.TabPage
	portGroupBox            *walk.GroupBox
	baudLabel               *walk.Label
	baudComboBox            *walk.ComboBox
	portComboBox            *walk.ComboBox
	portLabel               *walk.Label
	pushButton              *walk.PushButton
	portCheckBox            *walk.CheckBox
	parityGroupBox          *walk.GroupBox
	parityNoneRadioButton_2 *walk.RadioButton
	parityOddRadioButton    *walk.RadioButton
	parityEvenRadioButton   *walk.RadioButton
	parityMarkRadioButton   *walk.RadioButton
	paritySpaceRadioButton  *walk.RadioButton
	dataBitsGroupBox        *walk.GroupBox
	dataBits8RadioButton    *walk.RadioButton
	dataBits7RadioButton    *walk.RadioButton
	dataBits6RadioButton    *walk.RadioButton
	dataBits5RadioButton    *walk.RadioButton
	hwFlowCtrlGroupBox      *walk.GroupBox
	stopBitsGroupBox        *walk.GroupBox
	tab_3                   *walk.TabPage
	tab_4                   *walk.TabPage
	tab_5                   *walk.TabPage
	tab_6                   *walk.TabPage
	tab_7                   *walk.TabPage
	tab_8                   *walk.TabPage
	tab_9                   *walk.TabPage
	tab_10                  *walk.TabPage
	widget                  *walk.Composite
	statusGroupBox          *walk.GroupBox
}

func (w *Dialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	var font *walk.Font
	if font == nil {
		font = nil
	}

	w.SetName("Dialog")
	if err := w.SetClientSize(walk.Size{670, 628}); err != nil {
		return err
	}
	if err := w.SetTitle(`Dialog`); err != nil {
		return err
	}

	// textEdit
	if w.ui.textEdit, err = walk.NewTextEdit(w); err != nil {
		return err
	}
	w.ui.textEdit.SetName("textEdit")
	if err := w.ui.textEdit.SetBounds(walk.Rectangle{0, 0, 800, 300}); err != nil {
		return err
	}

	// tabWidget
	if w.ui.tabWidget, err = walk.NewTabWidget(w); err != nil {
		return err
	}
	w.ui.tabWidget.SetName("tabWidget")
	if err := w.ui.tabWidget.SetBounds(walk.Rectangle{0, 300, 571, 291}); err != nil {
		return err
	}

	// displayTab
	if w.ui.displayTab, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.displayTab.SetName("displayTab")
	if err := w.ui.displayTab.SetTitle(`Display`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.displayTab); err != nil {
		return err
	}

	// displayAsGroupBox
	if w.ui.displayAsGroupBox, err = walk.NewGroupBox(w.ui.displayTab); err != nil {
		return err
	}
	w.ui.displayAsGroupBox.SetName("displayAsGroupBox")
	if err := w.ui.displayAsGroupBox.SetBounds(walk.Rectangle{0, 0, 131, 271}); err != nil {
		return err
	}
	if err := w.ui.displayAsGroupBox.SetTitle(`Display As`); err != nil {
		return err
	}

	// portTab
	if w.ui.portTab, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.portTab.SetName("portTab")
	if err := w.ui.portTab.SetTitle(`Port`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.portTab); err != nil {
		return err
	}

	// portGroupBox
	if w.ui.portGroupBox, err = walk.NewGroupBox(w.ui.portTab); err != nil {
		return err
	}
	w.ui.portGroupBox.SetName("portGroupBox")
	if err := w.ui.portGroupBox.SetBounds(walk.Rectangle{0, 0, 561, 221}); err != nil {
		return err
	}
	if err := w.ui.portGroupBox.SetTitle(``); err != nil {
		return err
	}

	// baudLabel
	if w.ui.baudLabel, err = walk.NewLabel(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.baudLabel.SetName("baudLabel")
	if err := w.ui.baudLabel.SetBounds(walk.Rectangle{10, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.baudLabel.SetText(`Baud`); err != nil {
		return err
	}

	// baudComboBox
	if w.ui.baudComboBox, err = walk.NewComboBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.baudComboBox.SetName("baudComboBox")
	if err := w.ui.baudComboBox.SetBounds(walk.Rectangle{40, 20, 74, 22}); err != nil {
		return err
	}

	// portComboBox
	if w.ui.portComboBox, err = walk.NewComboBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portComboBox.SetName("portComboBox")
	if err := w.ui.portComboBox.SetBounds(walk.Rectangle{160, 20, 74, 22}); err != nil {
		return err
	}

	// portLabel
	if w.ui.portLabel, err = walk.NewLabel(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portLabel.SetName("portLabel")
	if err := w.ui.portLabel.SetBounds(walk.Rectangle{130, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.portLabel.SetText(`Port`); err != nil {
		return err
	}

	// pushButton
	if w.ui.pushButton, err = walk.NewPushButton(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.pushButton.SetName("pushButton")
	if err := w.ui.pushButton.SetBounds(walk.Rectangle{250, 20, 75, 23}); err != nil {
		return err
	}
	if err := w.ui.pushButton.SetText(`Open`); err != nil {
		return err
	}

	// portCheckBox
	if w.ui.portCheckBox, err = walk.NewCheckBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portCheckBox.SetName("portCheckBox")
	if err := w.ui.portCheckBox.SetBounds(walk.Rectangle{340, 20, 13, 13}); err != nil {
		return err
	}
	if err := w.ui.portCheckBox.SetText(``); err != nil {
		return err
	}

	// parityGroupBox
	if w.ui.parityGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.parityGroupBox.SetName("parityGroupBox")
	if err := w.ui.parityGroupBox.SetBounds(walk.Rectangle{10, 70, 73, 143}); err != nil {
		return err
	}
	if err := w.ui.parityGroupBox.SetTitle(`Parity`); err != nil {
		return err
	}

	// parityNoneRadioButton_2
	if w.ui.parityNoneRadioButton_2, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityNoneRadioButton_2.SetName("parityNoneRadioButton_2")
	if err := w.ui.parityNoneRadioButton_2.SetBounds(walk.Rectangle{11, 23, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityNoneRadioButton_2.SetText(`None`); err != nil {
		return err
	}

	// parityOddRadioButton
	if w.ui.parityOddRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityOddRadioButton.SetName("parityOddRadioButton")
	if err := w.ui.parityOddRadioButton.SetBounds(walk.Rectangle{11, 46, 41, 17}); err != nil {
		return err
	}
	if err := w.ui.parityOddRadioButton.SetText(`Odd`); err != nil {
		return err
	}

	// parityEvenRadioButton
	if w.ui.parityEvenRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityEvenRadioButton.SetName("parityEvenRadioButton")
	if err := w.ui.parityEvenRadioButton.SetBounds(walk.Rectangle{11, 69, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityEvenRadioButton.SetText(`Even`); err != nil {
		return err
	}

	// parityMarkRadioButton
	if w.ui.parityMarkRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityMarkRadioButton.SetName("parityMarkRadioButton")
	if err := w.ui.parityMarkRadioButton.SetBounds(walk.Rectangle{11, 92, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityMarkRadioButton.SetText(`Mark`); err != nil {
		return err
	}

	// paritySpaceRadioButton
	if w.ui.paritySpaceRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.paritySpaceRadioButton.SetName("paritySpaceRadioButton")
	if err := w.ui.paritySpaceRadioButton.SetBounds(walk.Rectangle{11, 115, 53, 17}); err != nil {
		return err
	}
	if err := w.ui.paritySpaceRadioButton.SetText(`Space`); err != nil {
		return err
	}

	// dataBitsGroupBox
	if w.ui.dataBitsGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.dataBitsGroupBox.SetName("dataBitsGroupBox")
	if err := w.ui.dataBitsGroupBox.SetBounds(walk.Rectangle{87, 70, 79, 143}); err != nil {
		return err
	}
	if err := w.ui.dataBitsGroupBox.SetTitle(`Data Bits`); err != nil {
		return err
	}

	// dataBits8RadioButton
	if w.ui.dataBits8RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits8RadioButton.SetName("dataBits8RadioButton")
	if err := w.ui.dataBits8RadioButton.SetBounds(walk.Rectangle{11, 27, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits8RadioButton.SetText(`8 bits`); err != nil {
		return err
	}

	// dataBits7RadioButton
	if w.ui.dataBits7RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits7RadioButton.SetName("dataBits7RadioButton")
	if err := w.ui.dataBits7RadioButton.SetBounds(walk.Rectangle{11, 54, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits7RadioButton.SetText(`7 bits`); err != nil {
		return err
	}

	// dataBits6RadioButton
	if w.ui.dataBits6RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits6RadioButton.SetName("dataBits6RadioButton")
	if err := w.ui.dataBits6RadioButton.SetBounds(walk.Rectangle{11, 81, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits6RadioButton.SetText(`6 bits`); err != nil {
		return err
	}

	// dataBits5RadioButton
	if w.ui.dataBits5RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits5RadioButton.SetName("dataBits5RadioButton")
	if err := w.ui.dataBits5RadioButton.SetBounds(walk.Rectangle{11, 108, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits5RadioButton.SetText(`5 bits`); err != nil {
		return err
	}

	// hwFlowCtrlGroupBox
	if w.ui.hwFlowCtrlGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.hwFlowCtrlGroupBox.SetName("hwFlowCtrlGroupBox")
	if err := w.ui.hwFlowCtrlGroupBox.SetBounds(walk.Rectangle{170, 126, 170, 80}); err != nil {
		return err
	}
	if err := w.ui.hwFlowCtrlGroupBox.SetTitle(`Hardware Flow Control`); err != nil {
		return err
	}

	// stopBitsGroupBox
	if w.ui.stopBitsGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.stopBitsGroupBox.SetName("stopBitsGroupBox")
	if err := w.ui.stopBitsGroupBox.SetBounds(walk.Rectangle{170, 70, 146, 51}); err != nil {
		return err
	}
	if err := w.ui.stopBitsGroupBox.SetTitle(`Stop Bits`); err != nil {
		return err
	}

	// tab_3
	if w.ui.tab_3, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_3.SetName("tab_3")
	if err := w.ui.tab_3.SetTitle(`Capture`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_3); err != nil {
		return err
	}

	// tab_4
	if w.ui.tab_4, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_4.SetName("tab_4")
	if err := w.ui.tab_4.SetTitle(`Pins`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_4); err != nil {
		return err
	}

	// tab_5
	if w.ui.tab_5, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_5.SetName("tab_5")
	if err := w.ui.tab_5.SetTitle(`Send`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_5); err != nil {
		return err
	}

	// tab_6
	if w.ui.tab_6, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_6.SetName("tab_6")
	if err := w.ui.tab_6.SetTitle(`Echo Port`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_6); err != nil {
		return err
	}

	// tab_7
	if w.ui.tab_7, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_7.SetName("tab_7")
	if err := w.ui.tab_7.SetTitle(`I2C`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_7); err != nil {
		return err
	}

	// tab_8
	if w.ui.tab_8, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_8.SetName("tab_8")
	if err := w.ui.tab_8.SetTitle(`I2C-2`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_8); err != nil {
		return err
	}

	// tab_9
	if w.ui.tab_9, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_9.SetName("tab_9")
	if err := w.ui.tab_9.SetTitle(`I2CMisc`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_9); err != nil {
		return err
	}

	// tab_10
	if w.ui.tab_10, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_10.SetName("tab_10")
	if err := w.ui.tab_10.SetTitle(`Misc`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_10); err != nil {
		return err
	}

	// widget
	if w.ui.widget, err = walk.NewComposite(w); err != nil {
		return err
	}
	w.ui.widget.SetName("widget")
	if err := w.ui.widget.SetBounds(walk.Rectangle{0, 600, 800, 32}); err != nil {
		return err
	}

	// statusGroupBox
	if w.ui.statusGroupBox, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.statusGroupBox.SetName("statusGroupBox")
	if err := w.ui.statusGroupBox.SetBounds(walk.Rectangle{580, 320, 81, 271}); err != nil {
		return err
	}
	if err := w.ui.statusGroupBox.SetTitle(`Status`); err != nil {
		return err
	}

	// Tab order

	succeeded = true

	return nil
}
